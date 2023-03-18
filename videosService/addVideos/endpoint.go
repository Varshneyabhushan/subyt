package addVideos

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type addVideosResponse struct {
	Count int
}

func MakeAddVideosEndpoint(s VideosAddingService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var videos []Video
		err := json.NewDecoder(r.Body).Decode(&videos)
		if err != nil {
			http.Error(w, "invalid body sent", http.StatusBadRequest)
			return
		}

		count, err := s.AddVideos(videos)
		if err != nil {
			http.Error(w, "error while adding videos : "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(addVideosResponse{count})
		if err != nil {
			log.Fatal("error while sending response in addVideos api : ", err)
		}
	}
}
