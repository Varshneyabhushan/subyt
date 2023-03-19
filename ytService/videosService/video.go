package videosservice

import "time"

type Video struct {
	Title         string
	Description   string
	Thumbnails    []Thumbnail
	Channel       Channel
	CreatedAt     time.Time
	VideoComparor `json:",inline"`
}

type Thumbnail struct {
	Width  int64
	Height int64
	Url    string
}

type Channel struct {
	Id    string
	Title string
}
