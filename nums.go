package main

import "fmt"

func main() {

	var a, b float64
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Println("Деление на ноль невозможно")
	}

	if b != 0 {
		fmt.Printf("%d %% %d = %d\n", int(a), int(b), int(a)%int(b))
	} else {
		fmt.Println("Остаток от деления на ноль невозможен")
	}

	a++
	b--
	fmt.Printf("После инкремента: a = %.2f\n", a)
	fmt.Printf("После декремента: b = %.2f\n", b)
}
