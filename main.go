package main

import (
	"diary_api/database"
	"diary_api/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDB()
}

func loadDB() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("gagal membaca file .env")
	}
}
