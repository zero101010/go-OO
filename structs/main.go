package main

import "fmt"
// criar um tipo derivado de um tipo primitivo
type SuperNumber int
func UseSuperNumber(){
	x := [] SuperNumber{1,3,4}
	fmt.Println(x)
}
// Struct car
type Car struct {
	Name string
	Year int
	Color string
}
// herança
type SuperCar struct {
	Car
	CanFly bool
}
func (c Car ) info() string {
	return c.Name
}

func main() {
	UseSuperNumber()
	car1 := Car{"GOl",2018,"Azul"}
	sCar := SuperCar{Car: Car{
		"Fusca",
		2030,
		"Blue",
	},CanFly: true}
	fmt.Println(sCar.info())
	fmt.Println(car1.info())
}