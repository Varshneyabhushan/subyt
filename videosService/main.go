package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"videosservice/env"
	"videosservice/httpServer"
	"videosservice/storageServices/elasticsearch"
	"videosservice/storageServices/mongo"
)

func main() {

	envConfig, err := env.LoadEnv()
	if err != nil {
		log.Fatal("error while reading .env", err)
		return
	}
	log.Println("environmental variables loaded successfully")

	mongoService, err := mongo.NewService(envConfig.MongoConfig)
	if err != nil {
		log.Fatal("error while connecting to database : ", err)
		return
	}
	log.Println("database connection established")

	if err := mongo.InitIndexes(mongoService); err != nil {
		log.Fatal("error while creating mongo indexes : ", err)
		return
	}
	log.Println("mongo indexes are created")

	esService, err := elasticsearch.NewService(envConfig.ElasticSearchConfig)
	if err != nil {
		log.Fatal("error while getting esClinet : ", err)
		return
	}
	log.Println("es connection established")

	router := httpServer.MakeRouter(mongoService, esService)
	address := fmt.Sprintf(":%s", strconv.Itoa(envConfig.ServerConfig.Port))
	log.Fatal(http.ListenAndServe(address, router))
}
