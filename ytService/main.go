package main

import (
	"context"
	"log"
	"time"
	"ytservice/env"
	"ytservice/limitidentifier"
	"ytservice/periodicjob"
	"ytservice/videofetcher"
	"ytservice/videosservice"
)

func main() {
        envConfig, err := env.GetConfig()
        if err != nil {
                log.Println("error while getting env : " + err.Error())
                return
        }

        fetcher, err := videofetcher.New(context.Background(), envConfig.YoutubeSearch)
        if err != nil {
                log.Fatal("error while making videoFetcher", err)
                return
        }

        videosService := videosservice.NewVideoService("http://videoservice")

        videoLimitIdentifier, err := limitidentifier.New("limits.json")
        if err != nil {
                log.Fatal("error while loading limit identifier", err)
                return
        }

        periodicJob := periodicjob.New(
                func() {
                        err := AddVideos(&videoLimitIdentifier, &fetcher, videosService, envConfig.YoutubeSearch.RequestCoolDown)
                        if err != nil {
                                log.Fatal("error while adding videos : " + err.Error())
                        }
                },
                time.Minute,
        )

        periodicJob.Start()
        time.Sleep(time.Minute)
}