package main

import (
	"fmt"
	"os"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	os.Remove("./example01.db")
	db, err := gorm.Open(sqlite.Open("./example01.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	u := User{Name: "Ivan"}
	res := db.Create(&u)
	fmt.Printf("User ID: %d, rows: %d\n", u.ID, res.RowsAffected)

	users := []User{{Name: "Petr1"}, {Name: "Maria1"}}

	for _, i := range users {
		db.Create(&i)
	}
	// batchSize - размер вставки за раз(insert many для одной сессию)
	db.CreateInBatches(users, 10)
}
