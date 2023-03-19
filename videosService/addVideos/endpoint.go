package addVideos

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"videosservice/videos"
)

type response struct {
	Count int
}

func MakeEndpoint(s Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var addingVideos []videos.Video
		err := json.NewDecoder(r.Body).Decode(&addingVideos)
		if err != nil {
			http.Error(w, "invalid body sent", http.StatusBadRequest)
			return
		}

		count, err := s.Add(addingVideos)
		if err != nil {
			http.Error(w, "error while adding videos : "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(response{count})
		if err != nil {
			log.Fatal("error while sending response in addVideos api : ", err)
		}
	}
}
