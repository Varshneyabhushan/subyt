package main

import (
	"context"
	"log"
	"time"
	"ytservice/env"
	"fmt"
	"ytservice/videofetcher"
)

func main() {
        envConfig, err := env.GetConfig()
        if err != nil {
                log.Println("error while getting env : " + err.Error())
                return
        }

        fetcher, err := videofetcher.New(context.Background(), envConfig.YoutubeSearch)
        if err != nil {
                fmt.Println("here we go", err)
        }

        for fetcher.HasNext() {
                timeout := time.After(5*time.Second)
                videos, err := fetcher.GetNext()
                fmt.Println("fetched videos of length", len(videos), err)
                for _, video := range videos {
                        fmt.Println(video.Title)
                }

                <- timeout
        }
}