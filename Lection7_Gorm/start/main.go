package main

import (
	"fmt"
	"os"

	// "gorm.io/driver/sqlite" // Драйверы(диалекты) конкретной СУБД
	"github.com/glebarez/sqlite" // without cgo
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"id,pk"`
	Name  string
	Email string
}

func main() {
	os.Remove("./test.db")

	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	u := User{Name: "User", Email: "user@gmail.com"}
	db.Create(&u)

	var recovered User
	db.First(&recovered, "name=?", "User")
	fmt.Println("Recovered", recovered)

	db.Model(&recovered).Update("Email", "newemail@mail.com")
	db.First(&recovered, 1)
	fmt.Println("After update", recovered)

	// db.Delete(&recovered, 1)
}
