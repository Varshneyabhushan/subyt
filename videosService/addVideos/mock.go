package addVideos

import (
	"fmt"
	"io"
)

type MockAddVideosService struct {
	writer io.Writer
}

func NewMockAddVideosService(w io.Writer) MockAddVideosService {
	return MockAddVideosService{writer: w}
}

func (s MockAddVideosService) AddVideos(videos []Video) (int, error) {
	for _, video := range videos {
		fmt.Fprintln(s.writer, video.Id, video.Title, video.PublishedAt)
	}

	return len(videos), nil
}
