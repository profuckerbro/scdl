package api

import (
	"strconv"
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
	var t Track
	//vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	fmt.Println(t.Url)
	scdlDownload(t.Url)

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(soundcloud.Filepath(t.Url)))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, soundcloud.Filepath(t.Url))
	

}

func scdlDownload(url string) {

	soundcloud.ExtractSong(url)

}
