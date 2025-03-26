package seeders

import (
	"digitaltrader/db"
	"digitaltrader/models"
	"fmt"
)

func SeedDB() {
	db := db.InitDB()
	fmt.Println("============")
	fmt.Printf("Seeding data into %v Database", db.Migrator().CurrentDatabase())
	fmt.Println("")

	if err := models.SeedPayments(); err != nil { // Seed the payments table
		panic("Failed to seed payments: " + err.Error())
	}

	// Migrate models
	models.MigratePage()                       // Migrate the Page model
	if err := models.SeedPages(); err != nil { // Seed the pages table
		panic("Failed to seed pages: " + err.Error())
	}

	fmt.Println("============")
	fmt.Printf("Models Fired into %v Database Successfully!", db.Migrator().CurrentDatabase())
	fmt.Println("")
	fmt.Println("============")
}
