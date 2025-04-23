package main

import (
	"log"
	"net/http"
	"user-service/db"
	"user-service/redis"
	"user-service/routes"
)

func main() {
	db.InitDB()
	redis.InitRedis()
	router := routes.SetupRouter()
	log.Println("User service running on port :8001")
	log.Fatal(http.ListenAndServe("8001", router))
}
