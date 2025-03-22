package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Wrapper struct {
	Arr [4]byte
}

func main() {
	wrap := Wrapper{
		Arr: [4]byte{11, 22, 33, 44},
	}

	fmt.Println("wrapper reflect ptr \t", uintptr(unsafe.Pointer(&wrap)))

	val := reflect.ValueOf(&wrap)
	val = val.Elem()

	fmt.Println("wrapper normal ptr \t", val.UnsafeAddr())

	arrayFill(val, []byte{1, 2, 3})
	fmt.Println(wrap)
}

func arrayFill(dst reflect.Value, data []byte) {
	arr := dst.FieldByName("Arr")

	fmt.Println("array reflect ptr \t", arr.UnsafeAddr())
	// size := arr.Len()
	// sval := arr.Slice(0, size)
	reflect.Copy(arr, reflect.ValueOf(data))
}

func arrayToSlice(arr reflect.Value) (int, reflect.Value) {
	ptr := reflect.New(arr.Type()).Elem()
	ptr.Set(arr)

	size := ptr.Len()
	sval := ptr.Slice(0, size)

	return size, sval
}
