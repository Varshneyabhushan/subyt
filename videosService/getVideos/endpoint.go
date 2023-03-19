package getVideos

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"videosservice/videos"
)

type response struct {
	videos []videos.Video
}

func MakeEndpoint(s Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		query := r.URL.Query()
		skipString := query.Get("skip")
		limitString := query.Get("limit")

		skip, _ := strconv.Atoi(skipString)
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			limit = 10
		}

		videosResult, err := s.Get(skip, limit)
		if err != nil {
			http.Error(w, "error while getting videosResult : "+err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(response{videos: videosResult})
		if err != nil {
			log.Fatal("error while sending response in getVideos api : ", err)
		}
	}
}
