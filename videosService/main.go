package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"videosservice/env"
	"videosservice/httpServer"
	"videosservice/repository"
	"videosservice/repository/mongo"
	"videosservice/storage"
)

func main() {

	envConfig, err := env.LoadEnv()
	if err != nil {
		log.Fatal("error while reading .env", err)
		return
	}
	log.Println("environmental variables loaded successfully")

	database, err := storage.GetDatabase(envConfig.MongoConfig)
	if err != nil {
		log.Fatal("error while connecting to database : ", err)
		return
	}
	log.Println("database connection established")

	videoRepository := mongo.NewRepository(database.Collection("videos"))
	videosRepo := repository.NewCompositeRepository(videoRepository)

	router := httpServer.MakeRouter(videosRepo)
	address := fmt.Sprintf(":%s", strconv.Itoa(envConfig.ServerConfig.Port))
	log.Fatal(http.ListenAndServe(address, router))
}
