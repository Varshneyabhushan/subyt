package httpServer

import (
	"github.com/julienschmidt/httprouter"
	"videosservice/addVideos"
	"videosservice/getVideos"
	"videosservice/repository"
	"videosservice/searchVideos"
)

func MakeRouter(repository repository.Repository) *httprouter.Router {
	router := httprouter.New()

	//adding routes
	router.POST("/", addVideos.MakeEndpoint(repository))
	router.GET("/", getVideos.MakeEndpoint(repository))
	router.GET("/search", searchVideos.MakeEndpoint(repository))

	return router
}
