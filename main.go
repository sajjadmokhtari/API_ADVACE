package main

import (
	"Web/data"
	"Web/data/db"
	"Web/router"
	"log"
)

func main() {

	if err := db.InitDb(); err != nil {
		log.Fatal("❌ Init failed:", err)
	}
	data.Up_1()

	if err := db.InitRedis(); err != nil {
		log.Fatal("❌ Init failed redis:", err)
	}

	r := router.SetUpRouter()

	r.Static("/front", "./front")

	r.Run(":8080")
}
