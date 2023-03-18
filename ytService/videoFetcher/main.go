package videofetcher

import (
	"context"
	"errors"
	"ytservice/env"
	"ytservice/videosservice"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

//follows iterator pattern
type VideoFetcher struct {
	query string
	maxResults int
	NextPageToken string
	hasInitiated bool
	ytService *youtube.Service
}

func (fetcher VideoFetcher) HasNext() bool {
	return !fetcher.hasInitiated || len(fetcher.NextPageToken) != 0
}

func (fetcher *VideoFetcher) GetNext() ([]videosservice.Video, error){
	response, err := fetcher.ytService.Search.List([]string{ "id", "snippet"}).
		Q(fetcher.query).
		MaxResults(int64(fetcher.maxResults)).
		Order("date").
		PageToken(fetcher.NextPageToken).
		Do()

	//TODO: handle errors according to the requirement
	if err != nil {
		return []videosservice.Video{}, errors.New("error while getting results from youtube" + err.Error())
	}

	fetcher.NextPageToken = response.NextPageToken
	fetcher.hasInitiated = true
	return getVideos(response.Items), nil
}


func New(ctx context.Context, youtubeConfig env.YoutubeSearchConfig) (VideoFetcher, error) {
	service, err := youtube.NewService(ctx, option.WithAPIKey(youtubeConfig.DevKey))
	if err != nil {
		return VideoFetcher{}, err
	}

	return VideoFetcher{ 
		query : youtubeConfig.Query,
		ytService: service,
		maxResults : youtubeConfig.MaxResults,
	}, nil
}
