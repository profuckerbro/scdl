package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func RunServ() {
	//	cmd.Execute() this is normal execution
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/download", downloadHandler)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func main() {

}
