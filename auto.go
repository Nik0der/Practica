package main

import "fmt"

type Engine struct {
	Type    string
	Power   int
	Volume  float64
	IsTurbo bool
}

type Auto struct {
	Brand   string
	Model   string
	Year    int
	Mileage int
	Engine  Engine
}

func (c Auto) PrintInfo() {
	fmt.Printf("%s %s (%d год)\n", c.Brand, c.Model, c.Year)
	fmt.Printf("Пробег: %d км\n", c.Mileage)
	fmt.Printf("Двигатель: %s, %.1f л, %d л.с., Турбо: %t\n",
		c.Engine.Type, c.Engine.Volume, c.Engine.Power, c.Engine.IsTurbo)
}

func (c *Auto) AddMileage(km int) {
	c.Mileage += km
	fmt.Printf("Пробег увеличен на %d км. Новый пробег: %d км\n", km, c.Mileage)
}

func main() {

	engine := Engine{
		Type:    "Бензиновый",
		Power:   150,
		Volume:  2.0,
		IsTurbo: true,
	}

	car := Auto{
		Brand:   "Toyota",
		Model:   "Camry",
		Year:    2022,
		Mileage: 15000,
		Engine:  engine,
	}

	bmw := Auto{
		Brand: "BMW",
		Model: "X5",
		Year:  2023,
		Engine: Engine{
			Type:    "Дизельный",
			Power:   249,
			Volume:  3.0,
			IsTurbo: true,
		},
	}

	car.PrintInfo()
	car.AddMileage(500)
	fmt.Println()

	bmw.PrintInfo()
	bmw.AddMileage(1000)
}
