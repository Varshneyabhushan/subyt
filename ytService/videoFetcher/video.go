package videofetcher 

import (
	"time"
	"google.golang.org/api/youtube/v3"
)

type Video struct {
	Id string
	Title string
	Description string
	Thumbnails []thumbnail
	PublishedAt time.Time
	Channel channel
	CreatedAt time.Time
}

type thumbnail struct {
	Width int64
	Height int64
	Url string
}

type channel struct {
	Id string
	Title string
}

func getVideos(items []*youtube.SearchResult) []Video {
	var videos []Video 
	for _, item := range items {
		video := Video {
			Id: item.Id.VideoId,
			Title: item.Snippet.Title,
			Description: item.Snippet.Description,
			Channel: channel{ 
				Id: item.Snippet.ChannelId,
				Title: item.Snippet.ChannelTitle,
			},
			CreatedAt: time.Now(),
		}

		video.Thumbnails = getThumbnails(*item.Snippet)
		video.PublishedAt, _ = time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		videos = append(videos, video)
	}

	return videos
}

func getThumbnails(snippet youtube.SearchResultSnippet) (result []thumbnail){
	thumbnails := []*youtube.Thumbnail {
		snippet.Thumbnails.High,
		snippet.Thumbnails.Medium,
		snippet.Thumbnails.Default,
	}

	for _, tmb := range thumbnails {
		result = append(result, thumbnail {
			Width: tmb.Width,
			Height: tmb.Height,
			Url: tmb.Url,
		})
	}

	return result
}