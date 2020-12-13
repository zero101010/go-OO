package main

// nome do container bot-db-mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	Name string
	Age  int
}
type Cars struct {
	Name string
	Year int
}

func connectGorn() {
	dbConnection := "root:12345678@tcp(127.0.0.1:3307)/bot-db"
	db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err, "Unable to access the database!")
	}
	fmt.Println("Conectou")
	// aplica a migração do modelo que queremos
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Cars{})
	// Cria um user com a struct que queremos
	user := Users{Name: "haha", Age: 18}
	cars := Cars{"gol", 18}
	// Aplica dentro do banco com o create
	resultUsers := db.Create(&user)
	resultCars := db.Create(&cars)
	fmt.Println(resultUsers.RowsAffected)
	fmt.Println(resultCars.RowsAffected)
	defer sqlDB.Close()
}
func main() {
	connectGorn()
}
