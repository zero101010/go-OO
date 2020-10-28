package main

import "fmt"

type Names []interface{}

// método da interface name
func (n *Names) load() {
	*n = Names{"Igor","Allan","Daniel","Jessica"}

}

func (n Names) printNames(){
	fmt.Println(n)
}

func main()  {
	var name Names
	name.load()
	name.printNames()

}