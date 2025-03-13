package main

import (
	// "digitaltrader/docs"
	"digitaltrader/migrations"
	"digitaltrader/routes"
)

func main() {

	migrations.MigrateDB()
	r := routes.SetupRouter()

	r.Run(":8085")
}
