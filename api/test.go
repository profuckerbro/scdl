package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/profuckerbro/scdl/soundcloud"
	//	"github.com/gorilla/mux"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var t Track
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	scdlDownload(t.Url)
	w.Header().Add("Content-Disposition", "Attachment")
	w.Header().Set("Content-Type", "audio/mpeg3")
	fmt.Println(soundcloud.Filepath(t.Url))

	http.ServeFile(w, r, soundcloud.Filepath(t.Url))

}
