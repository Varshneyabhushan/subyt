package repository

import (
	"videosservice/repository/elasticsearch"
	"videosservice/repository/mongo"
)

func toMongoVideo(video Video) mongo.Video {
	return mongo.Video{
		YtId:       video.Id,
		Properties: video.Properties,
		Channel:    video.Channel,
		Thumbnails: video.Thumbnails,
	}
}

func toEsVideo(video Video) elasticsearch.Video {
	return elasticsearch.Video{
		Id:           video.Id,
		Title:        video.Title,
		Description:  video.Description,
		ChannelTitle: video.Channel.Title,
	}
}

func ToVideo(mongoVideo mongo.Video) Video {
	return Video{
		Id:         mongoVideo.YtId,
		Thumbnails: mongoVideo.Thumbnails,
		Channel:    mongoVideo.Channel,
		Properties: mongoVideo.Properties,
	}
}

func ToVideos(mongoVideos []mongo.Video) []Video {
	var result []Video
	for _, video := range mongoVideos {
		result = append(result, ToVideo(video))
	}

	return result
}
