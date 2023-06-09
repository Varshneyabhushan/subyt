package jobexecution

import (
	"log"
	"ytworker/videofetcher"
)

func fetchVideoResponse(fetcher videofetcher.VideoFetcher, tracker *DelayTracker,
	nextPageToken string) (videofetcher.VideosResponse, bool, error) {
	response, err := fetcher.GetNext(nextPageToken)
	if err != nil {

		//quota exceeded
		if response.Status == 403 {
			log.Println("quota exceeded for this api key")
			tracker.ProportionalDelay()
			return response, true, nil
		}

		//if status is 4xx, requests has to be stopped
		if response.Status >= 400 && response.Status < 410 {
			return response, false, err
		}

		//client exhausted or server error
		if response.Status == 429 || response.Status >= 500 {
			log.Println("google server error : ", response.Status, err)
			tracker.ExponentialBackOff()
			return response, true, nil
		}

		log.Println("error while getting videos : ", err)
		tracker.ProportionalDelay()
		return response, true, nil
	}

	return response, false, nil
}
