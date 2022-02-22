package main

import (
	"demo/app"
	"demo/platform/db"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	db.Migrate()
	app.InitialiseApp()
}
