package videofetcher 

import (
	"time"
	"google.golang.org/api/youtube/v3"
	"ytservice/videosservice"
)

func getVideos(items []*youtube.SearchResult) []videosservice.Video {
	var videos []videosservice.Video
	for _, item := range items {
		video := videosservice.Video {
			Id: item.Id.VideoId,
			Title: item.Snippet.Title,
			Description: item.Snippet.Description,
			Channel: videosservice.Channel { 
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

func getThumbnails(snippet youtube.SearchResultSnippet) (result []videosservice.Thumbnail){
	thumbnails := []*youtube.Thumbnail {
		snippet.Thumbnails.High,
		snippet.Thumbnails.Medium,
		snippet.Thumbnails.Default,
	}

	for _, tmb := range thumbnails {
		result = append(result, videosservice.Thumbnail {
			Width: tmb.Width,
			Height: tmb.Height,
			Url: tmb.Url,
		})
	}

	return result
}