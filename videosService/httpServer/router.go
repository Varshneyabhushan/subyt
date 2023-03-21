package httpServer

import (
	"github.com/julienschmidt/httprouter"
	"videosservice/addVideos"
	"videosservice/getVideos"
	"videosservice/searchVideos"

	"videosservice/storageServices/elasticsearch"
	"videosservice/storageServices/mongo"
)

func MakeRouter(mongoService mongo.Service,
	esService elasticsearch.Service) *httprouter.
	Router {
	router := httprouter.New()

	mongoVideosService := mongo.GetCollection[mongo.Video](mongoService, "videos")
	esVideosService := elasticsearch.GetIndex[elasticsearch.Video](esService, "videos")

	addingService := addVideos.MakeAddService(mongoVideosService, esVideosService)
	router.POST("/", addVideos.MakeEndpoint(addingService))

	gettingService := getVideos.MakeGetVideosService(mongoVideosService)
	router.GET("/", getVideos.MakeEndpoint(gettingService))

	searchingService := searchVideos.MakeSearchService(mongoVideosService, esVideosService)
	router.GET("/search", searchVideos.MakeEndpoint(searchingService))

	return router
}
