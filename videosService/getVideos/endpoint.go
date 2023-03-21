package getVideos

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"videosservice/addVideos"
)

type response struct {
	Videos []addVideos.VideoResponse
}

type Service func(skip, limit int64) ([]addVideos.VideoResponse, error)

func MakeEndpoint(get Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		query := r.URL.Query()
		skipString := query.Get("skip")
		limitString := query.Get("limit")

		skip, _ := strconv.Atoi(skipString)
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			limit = 10
		}

		videosResult, err := get(int64(skip), int64(limit))
		if err != nil {
			http.Error(w, "error while getting videosResult : "+err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(response{Videos: videosResult})
		if err != nil {
			log.Fatal("error while sending response in getVideos api : ", err)
		}
	}
}
