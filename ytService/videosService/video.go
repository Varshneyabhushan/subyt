package videosservice

import "time"

type Video struct {
	Id string
	Title string
	Description string
	Thumbnails []Thumbnail
	PublishedAt time.Time
	Channel Channel
	CreatedAt time.Time
}

type Thumbnail struct {
	Width int64
	Height int64
	Url string
}

type Channel struct {
	Id string
	Title string
}