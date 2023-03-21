package addVideos

import (
	"videosservice/repository/elasticsearch"
	"videosservice/repository/mongo"
)

type VideoResponse struct {
	Id               string
	Thumbnails       []mongo.Thumbnail
	Channel          mongo.Channel
	mongo.Properties `json:",inline"`
}

func ToMongoVideo(videoResponse VideoResponse) mongo.Video {
	return mongo.Video{
		YtId:       videoResponse.Id,
		Properties: videoResponse.Properties,
		Channel:    videoResponse.Channel,
		Thumbnails: videoResponse.Thumbnails,
	}
}

func ToESVideo(videoResponse VideoResponse) elasticsearch.Video {
	return elasticsearch.Video{
		Id:           videoResponse.Id,
		Title:        videoResponse.Title,
		Description:  videoResponse.Description,
		ChannelTitle: videoResponse.Channel.Title,
	}
}

func ToVideo(mongoVideo mongo.Video) VideoResponse {
	return VideoResponse{
		Id:         mongoVideo.YtId,
		Thumbnails: mongoVideo.Thumbnails,
		Channel:    mongoVideo.Channel,
		Properties: mongoVideo.Properties,
	}
}

func ToVideoResponses(mongoVideos []mongo.Video) []VideoResponse {
	var result []VideoResponse
	for _, video := range mongoVideos {
		result = append(result, ToVideo(video))
	}

	return result
}
