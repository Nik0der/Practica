package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Printf("int: %d байт\n", unsafe.Sizeof(int(0)))
	fmt.Printf("int8: %d байт\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int16: %d байт\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("int32: %d байт\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int64: %d байт\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("uint: %d байт\n", unsafe.Sizeof(uint(0)))
	fmt.Printf("uint8: %d байт\n", unsafe.Sizeof(uint8(0)))
	fmt.Printf("uint16: %d байт\n", unsafe.Sizeof(uint16(0)))
	fmt.Printf("uint32: %d байт\n", unsafe.Sizeof(uint32(0)))
	fmt.Printf("uint64: %d байт\n", unsafe.Sizeof(uint64(0)))

	fmt.Printf("float32: %d байт\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("float64: %d байт\n", unsafe.Sizeof(float64(0)))

	fmt.Printf("complex64: %d байт\n", unsafe.Sizeof(complex64(0)))
	fmt.Printf("complex128: %d байт\n", unsafe.Sizeof(complex128(0)))

	fmt.Printf("byte: %d байт\n", unsafe.Sizeof(byte(0)))
	str := "test"
	fmt.Printf("string (базовый размер): %d байт\n", unsafe.Sizeof(str))

	fmt.Printf("bool: %d байт\n", unsafe.Sizeof(false))
}
