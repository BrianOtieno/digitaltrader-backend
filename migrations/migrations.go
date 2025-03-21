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
	models.MigrateAddress()
	models.MigrateCategory()
	models.MigrateChatMessage()
	models.MigrateChatRoom()
	models.MigrateCity()
	models.MigrateFailedJob()
	models.MigrateFavourite()
	models.MigrateFlush()
	models.MigrateGeneral()
	models.MigrateLanguage()
	models.MigrateOTP()
	models.MigratePasswordResetToken()
	models.MigratePayment()
	models.MigrateTransaction()
	models.MigrateTransfer()
	models.MigrateWallet()

	fmt.Println("============")
	fmt.Printf("Models Fired into %v Database Successfully!", db.Migrator().CurrentDatabase())
	fmt.Println("")
	fmt.Println("============")
}
