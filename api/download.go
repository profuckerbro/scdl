package api

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/profuckerbro/scdl/soundcloud"
	"strconv"
	"net/url"
	"io/ioutil"

//	"github.com/gorilla/mux"
)
type Track struct {
    Url string
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var t Track

//VALIDATOR
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	
	resp, err := http.Get(t.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		
	}
	// response
	body , err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		
	}
//VALIDATOR


	fmt.Println("SHOULD BE HIDDEN")
	var filepath = soundcloud.Filepath(body)
	var filename = soundcloud.Filename(body) 


	u, err := url.ParseRequestURI(t.Url)
	if err != nil {
   		log.Panicln(err)
	}
	fmt.Println(u)
	
	scdlDownload(t.Url)

	w.Header().Add("Content-Disposition", "Attachment; filename=" + strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")

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
