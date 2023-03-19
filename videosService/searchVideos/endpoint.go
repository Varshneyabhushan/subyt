package searchVideos

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"videosservice/repository"
)

type response struct {
	Videos []repository.Video
}

func MakeEndpoint(s Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		query := r.URL.Query()
		term := query.Get("term")

		searchedVideos, err := s.Search(term)
		if err != nil {
			http.Error(w, "error while getting searchedVideos : "+err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(response{Videos: searchedVideos})
		if err != nil {
			log.Fatal("error while sending response in searchVideos api : ", err)
		}
	}
}
