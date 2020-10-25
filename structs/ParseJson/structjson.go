package main

import (
	"encoding/json"
	"fmt"
)
// Para mudar a visibilidade do atributo da classe deve se modificar a letra inicial
// Letras maiúsculas public
// Letras minusculas private
//type Car struct {
//	Name string `json:"name"`
//	Year int `json:"-"`
//	color string `json:"color"`
//}

type Car struct {
	Name string
	Year int
	color string
}

// transformar objeto em json
func StructToJson(car Car) string{
	result, _ := json.Marshal(car)
	return string(result)
}

func main()  {
	//car := Car{"Gol",2017,"Amarelo"}
	var car1 Car
	data := []byte(`{"Name":"go"","Year":2017,"Color":"Black"}`)
	// bind json para struct
	json.Unmarshal(data, &car1)
	fmt.Println(car1.Name)
	// tem que converter para string por conta da tabela ascii
	//fmt.Println(StructToJson(car))
}