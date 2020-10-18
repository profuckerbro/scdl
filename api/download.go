package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/profuckerbro/scdl/soundcloud"

//	"github.com/gorilla/mux"
)
type Track struct {
    Url string
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var t Track
	//vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	fmt.Println(t.Url)
	scdlDownload(t.Url)
	w.Header().Add("Content-Disposition", "Attachment")
	w.Header().Set("Content-Type", "audio/mpeg3")
	fmt.Println(soundcloud.Filepath(t.Url))
	http.ServeFile(w, r, soundcloud.Filepath(t.Url))
	

}

func scdlDownload(url string) {

	soundcloud.ExtractSong(url)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
