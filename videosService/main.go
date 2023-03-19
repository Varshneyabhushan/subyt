package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"videosservice/addVideos"
	"videosservice/env"
	"videosservice/getVideos"
	"videosservice/repository"
	"videosservice/repository/mongo"
	"videosservice/searchVideos"
	"videosservice/storage"
)

func main() {

	envConfig, err := env.GetConfig()
	if err != nil {
		log.Fatal("error while reading .env")
		return
	}

	router := httprouter.New()

	database, err := storage.GetDatabase("", "videosService")
	videosCollection := database.Collection("repository")

	videoRepository := mongo.NewRepository(videosCollection)

	videosRepo := repository.NewCompositeRepository(videoRepository)

	router.POST("/", addVideos.MakeEndpoint(videosRepo))
	router.GET("/", getVideos.MakeEndpoint(videosRepo))
	router.GET("/search", searchVideos.MakeEndpoint(videosRepo))

	err = http.ListenAndServe(":"+strconv.Itoa(envConfig.ServerConfig.Port), router)
	if err != nil {
		log.Fatal("error while starting the server")
		return
	}
}
