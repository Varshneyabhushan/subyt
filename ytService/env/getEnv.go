package env

import (
	"os"
	"strconv"
	"time"
)

func getNumber(key string) int {
	result, _ := strconv.Atoi(os.Getenv(key))
	return result
}

func getSeconds(key string) time.Duration {
	return time.Duration(getNumber(key)) * time.Second
}

func GetFromEnv() Config {
	return Config{
		YoutubeSearch: YoutubeSearchConfig{
			DevKey:     os.Getenv("YOUTUBE_SEARCH_DEV_KEY"),
			Query:      os.Getenv("YOUTUBE_SEARCH_QUERY"),
			MaxResults: getNumber("YOUTUBE_MAX_RESULTS"),
		},
		Scheduler: SchedulerConfig{
			SyncCoolDown:    getSeconds("SCHEDULER_SYNC_COOL_DOWN"),
			RequestCoolDown: getSeconds("SCHEDULER_REQUEST_COOL_DOWN"),
		},
		CheckPointPath:  "checkPoint.json",
		VideoServiceURL: os.Getenv("VIDEOS_SERVICE_URL"),
	}
}
