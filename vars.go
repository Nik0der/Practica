package main

import "fmt"

func main() {

	var a int = -42
	var b uint = 42
	var c int8 = -127
	var d uint8 = 255
	var e rune = 'Я'

	var f float32 = 3.14
	var g float64 = 2.71828

	var h complex64 = 1 + 2i
	var i complex128 = complex(3, 4)

	var j byte = 'A'
	var k string = "Привет, Go!"

	var l bool = true

	fmt.Printf("Целые числа: int=%d, uint=%d, int8=%d, uint8=%d, rune='%c'\n", a, b, c, d, e)
	fmt.Printf("Числа с плавающей точкой: float32=%.2f, float64=%.5f\n", f, g)
	fmt.Printf("Комплексные числа: complex64=%v, complex128=%v\n", h, i)
	fmt.Printf("Байт и строка: byte='%c', string=\"%s\"\n", j, k)
	fmt.Printf("Логическое значение: bool=%t\n", l)
}
