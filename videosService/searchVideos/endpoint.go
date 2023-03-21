package searchVideos

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"videosservice/repository"
)

type response struct {
	Videos []repository.Video
}

func MakeEndpoint(s Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		query := r.URL.Query()
		term := query.Get("term")

		skipString := query.Get("skip")
		limitString := query.Get("limit")

		skip, _ := strconv.Atoi(skipString)
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			limit = 10
		}

		searchedVideos, err := s.Search(term, skip, limit)
		if err != nil {
			http.Error(w, "error while getting searchedVideos : "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(response{Videos: searchedVideos})
		if err != nil {
			log.Fatal("error while sending response in searchVideos api : ", err)
		}
	}
}
