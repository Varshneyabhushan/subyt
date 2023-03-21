package repository

import (
	"videosservice/repository/mongo"
)

//TODO: add jsonTags to fields

type Video struct {
	Id               string
	Thumbnails       []mongo.Thumbnail
	Channel          mongo.Channel
	mongo.Properties `json:",inline"`
}

type Repository interface {
	Add(videos []Video) (int, error)
	Get(skip, limit int64) ([]Video, error)
	Search(term string, skip, limit int) ([]Video, error)
}
