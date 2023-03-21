package httpServer

import (
	"github.com/julienschmidt/httprouter"
	"videosservice/addVideos"
	"videosservice/getVideos"
	"videosservice/repository/elasticsearch"
	"videosservice/searchVideos"

	mongo2 "videosservice/repository/mongo"
	es2 "videosservice/storageServices/elasticsearch"
	"videosservice/storageServices/mongo"
)

func MakeRouter(mongoService mongo.Service,
	esService es2.Service) *httprouter.
	Router {
	router := httprouter.New()

	mongoVideosService := mongo.GetCollection[mongo2.Video](mongoService, "videos")
	esVideosService := es2.GetIndex[elasticsearch.Video](esService, "videos")

	addingService := addVideos.MakeAddService(mongoVideosService, esVideosService)
	router.POST("/", addVideos.MakeEndpoint(addingService))

	gettingService := getVideos.MakeGetVideosService(mongoVideosService)
	router.GET("/", getVideos.MakeEndpoint(gettingService))

	searchingService := searchVideos.MakeSearchService(mongoVideosService, esVideosService)
	router.GET("/search", searchVideos.MakeEndpoint(searchingService))

	return router
}
