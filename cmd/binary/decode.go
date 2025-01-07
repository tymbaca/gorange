package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

const _tag = "bin"

type decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *decoder {
	return &decoder{r: r}
}

func (d *decoder) Decode(obj any, order binary.ByteOrder) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Pointer {
		return fmt.Errorf("obj must be a pointer, got: %v", val.Kind())
	}

	if val.IsNil() {
		return fmt.Errorf("obj must be a pointer to initialized variable (not nil), got: %v", val.Kind())
	}

	// dereference
	val = val.Elem()
	return decode(val, d.r, order)
}

func decode(val reflect.Value, from io.Reader, order binary.ByteOrder) error {
	switch val.Kind() {
	case reflect.Int8:
		i, err := readInt[int8](from, order)
		if err != nil {
			return err
		}

		val.SetInt(int64(i))

	case reflect.Int16:
		i, err := readInt[int16](from, order)
		if err != nil {
			return err
		}

		val.SetInt(int64(i))

	case reflect.Int32:
		i, err := readInt[int32](from, order)
		if err != nil {
			return err
		}

		val.SetInt(int64(i))

	case reflect.Int64:
		i, err := readInt[int64](from, order)
		if err != nil {
			return err
		}

		val.SetInt(int64(i))

	case reflect.Uint8:
		i, err := readUint[uint8](from, order)
		if err != nil {
			return err
		}

		val.SetUint(uint64(i))

	case reflect.Uint16:
		i, err := readUint[uint16](from, order)
		if err != nil {
			return err
		}

		val.SetUint(uint64(i))

	case reflect.Uint32:
		i, err := readUint[uint32](from, order)
		if err != nil {
			return err
		}

		val.SetUint(uint64(i))

	case reflect.Uint64:
		i, err := readUint[uint64](from, order)
		if err != nil {
			return err
		}

		val.SetUint(uint64(i))

	case reflect.Slice:
		return fmt.Errorf("not implemented: %v", val.Kind())

	case reflect.String:
		return fmt.Errorf("not implemented: %v", val.Kind())

	case reflect.Struct:
		for i := range val.NumField() {
			fieldVal := val.Field(i)
			err := decode(fieldVal, from, order)
			if err != nil {
				return fmt.Errorf("can't decode %v (%v): %w", fieldVal.Kind(), fieldVal, err)
			}
		}

	default:
		fmt.Println("ignoring field:", val)
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
