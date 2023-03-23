package main

import (
	"context"
	"log"
	"ytworker/env"
	"ytworker/jobexecution"
	"ytworker/storage"
	"ytworker/videofetcher"
	"ytworker/videosservice"
)

func main() {
	envConfig, err := env.LoadEnv()
	if err != nil {
		log.Println("error while getting env : " + err.Error())
		return
	}

	checkPoint := storage.LoadCheckPoint(envConfig.CheckPointPath)
	fetcher, err := videofetcher.New(context.Background(), envConfig.YoutubeSearch)
	if err != nil {
		log.Fatal("error while making videoFetcher", err)
		return
	}

	videosService := videosservice.NewVideoService(envConfig.VideoServiceURL)

	syncVideoJob := jobexecution.MakeSyncVideosJob(checkPoint, fetcher, videosService, envConfig.Scheduler)
	log.Fatal(syncVideoJob.Execute())
}
