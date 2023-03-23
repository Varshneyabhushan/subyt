package getVideosCount

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Service func() (int64, error)

type Response struct {
	Count int64
}

func MakeEndpoint(getCount Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		count, err := getCount()
		if err != nil {
			http.Error(w, "error while getting count of videos : "+err.Error(),
				http.StatusBadRequest)
			return
		}

		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(Response{count})
		if err != nil {
			log.Fatal("error while sending response in addVideos api : ", err)
		}
	}
}
