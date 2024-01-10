package main

import (
	"blogApp/api/database"
	"blogApp/api/initializers"
	"blogApp/api/models"
	"log"
)

func init() {
	config, err := initializers.LoadEnv("../..")
	if err != nil {
		log.Fatal("have some problem about env in migrate.go")
	}
	database.ConnectDB(&config)
}

func main() {
	// database.GlobalDB.AutoMigrate(&models.Bir{}) // en az bağımlılık içeren en üstte olmalı
	// database.GlobalDB.AutoMigrate(&models.Cok{})
	database.GlobalDB.AutoMigrate(&models.User{}) // 1- n post
	database.GlobalDB.AutoMigrate(&models.Blog{}) // n -1 user // 1 - n
	database.GlobalDB.AutoMigrate(&models.Emote{})
	database.GlobalDB.AutoMigrate(&models.Comment{}) // 1 - n
	// emotes n-1
	database.GlobalDB.AutoMigrate(&models.Categories{})
	database.GlobalDB.AutoMigrate(&models.Catepostrel{})

}
