package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)


func RunServ() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/download", downloadHandler)
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))

}


