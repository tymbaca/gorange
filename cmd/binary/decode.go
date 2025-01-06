package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

type decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *decoder {
	return &decoder{r: r}
}

func (d *decoder) Decode(data any, order binary.ByteOrder) error {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Pointer {
		return fmt.Errorf("data must be a pointer, got: %v", val.Kind())
	}

	if val.IsNil() {
		return fmt.Errorf("data must be a pointer to initialized variable (not nil), got: %v", val.Kind())
	}

	// dereference
	val = val.Elem()

	for i := range val.NumField() {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.Int8:
			i, err := readInt[int8](d.r, order)
			if err != nil {
				return err
			}

			field.SetInt(int64(i))

		case reflect.Int16:
			i, err := readInt[int16](d.r, order)
			if err != nil {
				return err
			}

			field.SetInt(int64(i))

		case reflect.Int32:
			i, err := readInt[int32](d.r, order)
			if err != nil {
				return err
			}

			field.SetInt(int64(i))

		case reflect.Int64:
			i, err := readInt[int64](d.r, order)
			if err != nil {
				return err
			}

			field.SetInt(int64(i))

		case reflect.Uint8:
			i, err := readUint[uint8](d.r, order)
			if err != nil {
				return err
			}

			field.SetUint(uint64(i))

		case reflect.Uint16:
			i, err := readUint[uint16](d.r, order)
			if err != nil {
				return err
			}

			field.SetUint(uint64(i))

		case reflect.Uint32:
			i, err := readUint[uint32](d.r, order)
			if err != nil {
				return err
			}

			field.SetUint(uint64(i))

		case reflect.Uint64:
			i, err := readUint[uint64](d.r, order)
			if err != nil {
				return err
			}

			field.SetUint(uint64(i))

		case reflect.Slice:
			panic(errors.New("not implemented"))

		case reflect.String:
			panic(errors.New("not implemented"))

		default:
			fmt.Println("ignoring field:", field)
		}
	}

	return nil
}

func readInt[I int8 | int16 | int32 | int64](r io.Reader, order binary.ByteOrder) (I, error) {
	var i I
	if err := binary.Read(r, order, &i); err != nil {
		return i, fmt.Errorf("can't read int: %w", err)
	}

	return i, nil
}

func readUint[U uint8 | uint16 | uint32 | uint64](r io.Reader, order binary.ByteOrder) (U, error) {
	var u U
	if err := binary.Read(r, order, &u); err != nil {
		return u, fmt.Errorf("can't read uint: %w", err)
	}

	return u, nil
}
