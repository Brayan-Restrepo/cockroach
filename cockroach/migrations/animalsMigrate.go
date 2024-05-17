package main

import (
	"cockroach/cockroach/entities"
	"cockroach/config"
	"cockroach/database"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	animalsMigrate(db)
}

func animalsMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Animals{})
}
