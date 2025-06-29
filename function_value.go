package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func incrementByValue(num int) {
	num++
	fmt.Println("Внутри incrementByValue:", num)
}

func incrementByPointer(num *int) {
	*num++
	fmt.Println("Внутри incrementByPointer:", *num)
}

func changePersonByValue(p Person) {
	p.Age += 5
	fmt.Println("Внутри changePersonByValue:", p)
}

func changePersonByPointer(p *Person) {
	p.Age += 5
	fmt.Println("Внутри changePersonByPointer:", *p)
}

func main() {
	value := 10
	pointer := 10

	fmt.Println("Перед вызовами:")
	fmt.Println("value =", value, "pointer =", pointer)

	incrementByValue(value)
	fmt.Println("После incrementByValue:", value)

	incrementByPointer(&pointer)
	fmt.Println("После incrementByPointer:", pointer)

	p1 := Person{"Анна", 25}
	fmt.Println("\nДо изменения (по значению):", p1)
	changePersonByValue(p1)
	fmt.Println("После changePersonByValue:", p1)

	p2 := Person{"Иван", 30}
	fmt.Println("\nДо изменения (по указателю):", p2)
	changePersonByPointer(&p2)
	fmt.Println("После changePersonByPointer:", p2)
}
