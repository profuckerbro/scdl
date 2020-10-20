package api

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"net/url"
	"io/ioutil"
	"github.com/profuckerbro/scdl/soundcloud"
)
type Metadata struct {
    Url string
}

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	
	
	
	var t Metadata

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
	fmt.Println(body)
//VALIDATOR


imageUrl := soundcloud.GetArtworkUrl(body)
fmt.Println(imageUrl)


//RETURN DATA
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{Response: success, imageUrl: ` + imageUrl + `,}`))
w.WriteHeader(http.StatusOK)
return


	
	/*SET HEADER
	w.Header().Add("Content-Disposition", "Attachment; filename=" + strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	*/

	

}



