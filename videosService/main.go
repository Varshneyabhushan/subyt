package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"videosservice/addVideos"
	"videosservice/env"
	"videosservice/getVideos"
	"videosservice/searchVideos"
	"videosservice/videos"
)

func main() {

	envConfig, err := env.GetConfig()
	if err != nil {
		log.Fatal("error while reading .env")
		return
	}

	router := httprouter.New()

	videoRepo := videos.NewMockRepository()

	router.POST("/videos", addVideos.MakeEndpoint(videoRepo))
	router.GET("/videos", getVideos.MakeEndpoint(videoRepo))
	router.GET("/search", searchVideos.MakeEndpoint(videoRepo))

	err = http.ListenAndServe(":"+strconv.Itoa(envConfig.ServerConfig.Port), router)
	if err != nil {
		log.Fatal("error while starting the server")
		return
	}
}
