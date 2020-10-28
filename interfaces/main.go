package main

import "fmt"


type vehicle interface {
	start() string
}

type car struct {
	name string
}

func (c car) start() string {
	return "The car " + c.name +" has been start"
}

type motorcycle struct {
	name string
}



func (mc motorcycle) start() string {
	return "The motorcycle " + mc.name +" has been start"
}

func startVehicle ( v vehicle) string {
	return v.start()
}
func main()  {
	c := car{"Gol"}
	mc:= motorcycle{"moto2"}
	fmt.Println(startVehicle(mc))
	fmt.Println(startVehicle(c))
	//fmt.Println(c.start())
	//fmt.Println(mc.start())
}
