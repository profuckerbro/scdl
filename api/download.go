package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/profuckerbro/scdl/soundcloud"
	"strconv"
	"net/url"

//	"github.com/gorilla/mux"
)
type Track struct {
    Url string
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var t Track
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	var filepath = soundcloud.Filepath(t.Url)
	var filename = soundcloud.Filename(t.Url) 


	u, err := url.ParseRequestURI(t.Url)
if err != nil {
   panic(err)
}
	fmt.Println("This is u")
	fmt.Println(u)
	scdlDownload(t.Url)

	w.Header().Add("Content-Disposition", "Attachment; filename=" + strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	fmt.Println(soundcloud.Filepath(t.Url))
	http.ServeFile(w, r, filepath)
	

}

func scdlDownload(url string) {

	soundcloud.ExtractSong(url)

}

func enableCors(w *http.ResponseWriter) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Content-Disposition, Accept-Encoding, Authorization,X-CSRF-Token"
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	(*w).Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
}
