package main

import (
	"context"
	"log"
	"time"
	"ytservice/env"
	"ytservice/periodicjob"
	"ytservice/videofetcher"
	"ytservice/videosservice"
        "ytservice/storage"
)

func main() {
        envConfig, err := env.GetConfig()
        if err != nil {
                log.Println("error while getting env : " + err.Error())
                return
        }

        checkPoint, err := storage.LoadCheckPoint("checkPoint.json")
        if err != nil {
                log.Fatal("error while loading checkpoint", err)
        }

        fetcher, err := videofetcher.New(context.Background(), envConfig.YoutubeSearch)
        if err != nil {
                log.Fatal("error while making videoFetcher", err)
                return
        }

        videosService := videosservice.NewVideoService("http://videoservice")

        periodicJob := periodicjob.New(
                func() {
                        err := AddVideos(checkPoint, fetcher, videosService, envConfig.YoutubeSearch.RequestCoolDown)
                        if err != nil {
                                log.Fatal("error while adding videos : " + err.Error())
                        }
                },
                time.Minute,
        )

        periodicJob.Start()
        time.Sleep(time.Minute)
}