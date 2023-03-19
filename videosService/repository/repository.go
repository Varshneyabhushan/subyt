package repository

import (
	"videosservice/repository/mongo"
)

type Video struct {
	Id               string
	Thumbnails       []mongo.Thumbnail
	Channel          mongo.Channel
	mongo.Properties `json:",inline"`
}

func mongoVideo(video Video) mongo.Video {
	return mongo.Video{
		YtId:       video.Id,
		Properties: video.Properties,
		Channel:    video.Channel,
		Thumbnails: video.Thumbnails,
	}
}

func ToVideo(mongoVideo mongo.Video) Video {
	return Video{
		Id:         mongoVideo.YtId,
		Properties: mongoVideo.Properties,
		Channel:    mongoVideo.Channel,
		Thumbnails: mongoVideo.Thumbnails,
	}
}

type Repository interface {
	Add(videos []Video) (int, error)
	Get(skip, limit int64) ([]Video, error)
	Search(term string) ([]Video, error)
}
