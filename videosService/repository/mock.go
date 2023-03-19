package repository

import (
	"strings"
)

type MockRepository struct {
	videos map[string]Video
}

func NewMockRepository() *MockRepository {
	return &MockRepository{videos: make(map[string]Video, 0)}
}

func (repo *MockRepository) Add(videos []Video) (int, error) {
	totalVideosAdded := 0
	for _, video := range videos {
		if _, ok := repo.videos[video.Id]; !ok {
			totalVideosAdded += 1
			repo.videos[video.Id] = video
		}
	}

	return totalVideosAdded, nil
}

func (repo *MockRepository) Get(skip, limit int) ([]Video, error) {
	index := 0
	result := make([]Video, 0)
	for _, video := range repo.videos {
		if len(result) >= limit {
			break
		}

		if skip <= index {
			result = append(result, video)
		}

		index += 1
	}

	return result, nil
}

func (repo *MockRepository) Search(term string) ([]Video, error) {
	var result []Video
	for _, video := range repo.videos {
		if strings.Contains(video.Title, term) {
			result = append(result, video)
		}
	}

	return result, nil
}
