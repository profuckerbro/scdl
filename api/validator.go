package api


import (
	"net/http"
	"io/ioutil"
	"log"
)

func Validator (url string) []byte{
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	// response
	body , err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	 return body
}

