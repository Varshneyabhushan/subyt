package httpServer

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"videosservice/addVideos"
	"videosservice/getVideos"
	"videosservice/getVideosCount"
	"videosservice/searchVideos"

	"videosservice/storageServices/elasticsearch"
	"videosservice/storageServices/mongo"
)

func cors(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter,
		r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r, ps)
	}
}

func MakeRouter(mongoService mongo.Service,
	esService elasticsearch.Service) *httprouter.
	Router {
	router := httprouter.New()

	mongoVideosService := mongo.GetCollection[mongo.Video](mongoService, "videos")
	esVideosService := elasticsearch.GetIndex[elasticsearch.Video](esService, "videos")

	addingService := addVideos.MakeAddService(mongoVideosService, esVideosService)
	router.POST("/", addVideos.MakeEndpoint(addingService))

	gettingService := getVideos.MakeGetVideosService(mongoVideosService)
	router.GET("/", cors(getVideos.MakeEndpoint(gettingService)))

	videosCountService := getVideosCount.MakeService(mongoVideosService)
	router.GET("/count", cors(getVideosCount.MakeEndpoint(videosCountService)))

	searchingService := searchVideos.MakeSearchService(mongoVideosService, esVideosService)
	router.GET("/search", cors(searchVideos.MakeEndpoint(searchingService)))

	return router
}
