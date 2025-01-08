package sbinary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"unsafe"
)

func Marshal(data any, order binary.ByteOrder) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.Grow(int(unsafe.Sizeof(data)))

	if err := NewEncoder(buf).Encode(data, order); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *encoder {
	return &encoder{w: w}
}

func (e *encoder) Encode(data any, order binary.ByteOrder) error {
	val := reflect.ValueOf(data)

	return encode(val, e.w, order)
}

func encode(val reflect.Value, to io.Writer, order binary.ByteOrder) error {
	buf := bytes.NewBuffer(nil)
	buf.Grow(int(val.Type().Size()))

	switch val.Kind() {
	case reflect.Int8:
		writeInt[int8](buf, order, val)

	case reflect.Int16:
		writeInt[int16](buf, order, val)

	case reflect.Int32:
		writeInt[int32](buf, order, val)

	case reflect.Int64:
		writeInt[int64](buf, order, val)

	case reflect.Uint8:
		writeUint[uint8](buf, order, val)

	case reflect.Uint16:
		writeUint[uint16](buf, order, val)

	case reflect.Uint32:
		writeUint[uint32](buf, order, val)

	case reflect.Uint64:
		writeUint[uint64](buf, order, val)

	case reflect.Slice:
		elemKind := val.Type().Elem().Kind()
		if elemKind == reflect.Uint8 {
			buf.Write(val.Bytes())
		}

	case reflect.String:
		buf.Write([]byte(val.String()))

	case reflect.Struct:
		for i := range val.NumField() {
			fieldVal := val.Field(i)
			err := encode(fieldVal, to, order)
			if err != nil {
				return fmt.Errorf("can't encode %v (%v): %w", fieldVal.Kind(), fieldVal, err)
			}
		}

	default:
		fmt.Println("ignoring field:", val)
	}

	_, err := io.Copy(to, buf)
	if err != nil {
		return fmt.Errorf("can't encode the struct: %w", err)
	}

	return nil
}

func writeInt[I int8 | int16 | int32 | int64](buf *bytes.Buffer, order binary.ByteOrder, val reflect.Value) {
	i := I(val.Int())
	if err := binary.Write(buf, order, i); err != nil {
		panic(err) // TODO:
	}
}

func writeUint[U uint8 | uint16 | uint32 | uint64](buf *bytes.Buffer, order binary.ByteOrder, val reflect.Value) {
	u := U(val.Uint())
	if err := binary.Write(buf, order, u); err != nil {
		panic(err)
	}
}
