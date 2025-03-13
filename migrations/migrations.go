package migrations

import (
	"digitaltrader/db"
	"digitaltrader/models"
	"fmt"
)

func MigrateDB() {
	db := db.InitDB()
	fmt.Println("============")
	fmt.Printf("Firing models into %v Database", db.Migrator().CurrentDatabase())
	fmt.Println("")

	models.MigrateActivityLog()
	models.MigrateUser()

	fmt.Println("============")
	fmt.Printf("Models Fired into %v Database Successfully!", db.Migrator().CurrentDatabase())
	fmt.Println("")
	fmt.Println("============")
}
