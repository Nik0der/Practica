package main

import "fmt"

type Transport interface {
	Move(speed int) string
	Stop() string
}

type Car struct {
	Model string
}

func (c Car) Move(speed int) string {
	return fmt.Sprintf("Автомобиль %s едет со скоростью %d км/ч", c.Model, speed)
}

func (c Car) Stop() string {
	return fmt.Sprintf("Автомобиль %s остановился", c.Model)
}

type Bicycle struct {
	Brand string
}

func (b Bicycle) Move(speed int) string {
	return fmt.Sprintf("Велосипед %s движется со скоростью %d км/ч", b.Brand, speed)
}

func (b Bicycle) Stop() string {
	return fmt.Sprintf("Велосипед %s остановлен", b.Brand)
}

func testTransport(t Transport, speed int) {
	fmt.Println(t.Move(speed))
	fmt.Println(t.Stop())
}

func main() {

	car := Car{Model: "Toyota Camry"}
	bike := Bicycle{Brand: "Stels"}

	testTransport(car, 60)
	testTransport(bike, 20)

	transports := []Transport{car, bike}
	for _, t := range transports {
		testTransport(t, 30)
	}
}
