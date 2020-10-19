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
)
type Track struct {
    Url string
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	
	
	
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
	u, err := url.ParseRequestURI(t.Url)
	if err != nil {
   		log.Panicln(err)
	}
	fmt.Println(u)
//VALIDATOR


	var filepath = soundcloud.Filepath(body)
	var filename = soundcloud.Filename(body) 


	
	soundcloud.ExtractSong(t.Url)

	//SET HEADER
	w.Header().Add("Content-Disposition", "Attachment; filename=" + strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	//SERVE FILE
	http.ServeFile(w, r, filepath)
	

}



