package soundcloud

import (
	"net/http"
	"io/ioutil"
	"log"
)

func Filepath (url string) string{
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	// response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	songname := GetTitle(body)
	artist := GetArtist(body)

	return artist + "/" + songname + ".mp3"
	 
}
func Filename (url string) string{
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	// response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	songname := GetTitle(body)
	

	return  songname + ".mp3"
	 
}