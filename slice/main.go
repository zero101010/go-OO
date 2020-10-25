package main

import "fmt"

func main()  {
	// o número 2 é a quantidade e o 5 é a capacidade que é duplicada toda vez que passa a quantidade
	slice := make([] int, 3,5)
	fmt.Println(cap(slice))
	fmt.Println(slice)
	// só é possível dar append em slice
	slice = append(slice,3,3,4,5,2)

	fmt.Println(slice)
	fmt.Println(cap(slice))

}
