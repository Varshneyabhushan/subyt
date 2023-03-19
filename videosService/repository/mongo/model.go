package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Video struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	YtId       string
	Properties `json:",inline" bson:",inline"`
	Thumbnails []Thumbnail
	Channel    Channel
}

type Properties struct {
	Title       string
	Description string
	PublishedAt time.Time
	CreatedAt   time.Time
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
