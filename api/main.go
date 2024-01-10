package main

import (
	"blogApp/api/database"
	"blogApp/api/initializers"
	"blogApp/api/route"
	"log"
)

func init() {
	config, err := initializers.LoadEnv("..")
	if err != nil {
		log.Fatal("some problems in  main")
	}
	database.ConnectDB(&config)
}

func main() {
	r := route.SetupRouter()
	r.Run(":8080")
}
