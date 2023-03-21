package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"videosservice/repository/elasticsearch"
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

func esVideo(Id primitive.ObjectID, video Video) elasticsearch.Video {
	return elasticsearch.Video{
		Id:          Id.Hex(),
		Title:       video.Title,
		Description: video.Description,
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

func ToVideos(mongoVideos []mongo.Video) []Video {
	var result []Video
	for _, video := range mongoVideos {
		result = append(result, ToVideo(video))
	}

	return result
}

type Repository interface {
	Add(videos []Video) (int, error)
	Get(skip, limit int64) ([]Video, error)
	Search(term string) ([]Video, error)
}
