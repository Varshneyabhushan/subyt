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

	envConfig, err := env.GetConfigFromFile("env.json")
	if err != nil {
		log.Fatal("error while reading .env")
		return
	}

	database, err := storage.GetDatabase(envConfig.MongoConfig)
	if err != nil {
		log.Fatal("error while connecting to database : ", err)
		return
	}

	videoRepository := mongo.NewRepository(database.Collection("videos"))
	videosRepo := repository.NewCompositeRepository(videoRepository)

	router := httpServer.MakeRouter(videosRepo)
	address := fmt.Sprintf(":%s", strconv.Itoa(envConfig.ServerConfig.Port))
	log.Fatal(http.ListenAndServe(address, router))
}
