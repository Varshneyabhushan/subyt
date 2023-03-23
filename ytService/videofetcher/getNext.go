package videofetcher

import (
	"context"
	"errors"
	"google.golang.org/api/googleapi"
	"ytworker/env"
	"ytworker/videosservice"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type VideoFetcher struct {
	query      string
	maxResults int
	ytService  *youtube.Service
}

func New(ctx context.Context, youtubeConfig env.YoutubeSearchConfig) (*VideoFetcher, error) {
	service, err := youtube.NewService(ctx, option.WithAPIKey(youtubeConfig.DevKey))
	if err != nil {
		return nil, err
	}

	return &VideoFetcher{
		query:      youtubeConfig.Query,
		ytService:  service,
		maxResults: youtubeConfig.MaxResults,
	}, nil
}

type VideosResponse struct {
	Status        int
	Videos        []videosservice.Video
	NextPageToken string
}

func toVideosResponse(response *youtube.SearchListResponse) VideosResponse {
	videos := parseVideos(response.Items)
	return VideosResponse{
		Status:        response.HTTPStatusCode,
		Videos:        videos,
		NextPageToken: response.NextPageToken,
	}
}

// GetNext gets the results from nextPage as long as nextPageToken is available
// else, starts getting new results
func (fetcher *VideoFetcher) GetNext(nextPageToken string) (VideosResponse, error) {
	response, err := fetcher.ytService.Search.List([]string{"id", "snippet"}).
		Q(fetcher.query).
		MaxResults(int64(fetcher.maxResults)).
		Order("date").
		PageToken(nextPageToken).
		Do()

	//TODO: handle errors according to the requirement
	if err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			return VideosResponse{Status: e.Code}, e
		}

		return VideosResponse{}, errors.New("error while getting results from youtube : " + err.
			Error())
	}

	return toVideosResponse(response), nil
}
