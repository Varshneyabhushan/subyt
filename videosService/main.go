package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"videosservice/env"
)

func main() {

	envConfig, err := env.GetConfig()
	if err != nil {
		log.Fatal("error while reading .env")
		return
	}

	router := httprouter.New()

	//addVideosService := addVideos.NewMockAddVideosService(os.Stdout)

	//router.POST("/videos", addVideos.MakeAddVideosEndpoint(addVideosService))

	err = http.ListenAndServe(":"+strconv.Itoa(envConfig.ServerConfig.Port), router)
	if err != nil {
		log.Fatal("error while starting the server")
		return
	}
}
