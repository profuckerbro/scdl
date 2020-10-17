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
	var t Track
	//vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	fmt.Println(t.Url)
	scdlDownload(t.Url)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "DOWNLOAD ROUTE"}`))

}

func scdlDownload(url string) {

	soundcloud.ExtractSong(url)

}
