package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"videosservice/env"
	"videosservice/httpServer"
	"videosservice/repository"
	"videosservice/repository/elasticsearch"
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

	esClient, err := storage.GetESClient(envConfig.ElasticSearchConfig)
	if err != nil {
		log.Fatal("error while getting esClinet : ", err)
		return
	}
	log.Println("es connection established")

	videoRepository := mongo.NewRepository(database.Collection("videos"))
	esRepository := elasticsearch.NewRepository(esClient)

	videosRepo := repository.NewCompositeRepository(videoRepository, esRepository)

	router := httpServer.MakeRouter(videosRepo)
	address := fmt.Sprintf(":%s", strconv.Itoa(envConfig.ServerConfig.Port))
	log.Fatal(http.ListenAndServe(address, router))
}
