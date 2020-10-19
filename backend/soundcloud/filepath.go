package soundcloud

import (
//	"net/http"
//	"io/ioutil"
	//"log"
)

func Filepath (body []byte) string{
	
	songname := GetTitle(body)
	artist := GetArtist(body)

	return artist + "/" + songname + ".mp3"
	 
}
func Filename (body []byte) string{
	

	songname := GetTitle(body)
	

	return  songname + ".mp3"
	 
}
