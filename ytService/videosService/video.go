package videosservice

import "time"

type VideoComparor struct {
	Id string
	PublishedAt time.Time
}

type Video struct {
	Title string
	Description string
	Thumbnails []Thumbnail
	Channel Channel
	CreatedAt time.Time
	VideoComparor `json:"-"`
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

func (v VideoComparor) IsEmpty() bool {
	return len(v.Id) == 0
}

func LimitCrossed(v1, v2 VideoComparor) bool {
	if(v1.Id == v2.Id) {
		return true
	}

	if(v2.Id == "") {
		return false
	}

	return v1.PublishedAt.After(v2.PublishedAt)
}

func GetValidVideos(videos []Video, limit VideoComparor) int {
	result := 0
	if len(videos) == 0 {
		return result
	}

	for _, video := range videos {
		if LimitCrossed(video.VideoComparor, limit) {
			return result
		}

		result += 1
	}

	return result
}