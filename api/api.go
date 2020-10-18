package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)


func RunServ() {
	//	cmd.Execute() this is normal execution
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/download", downloadHandler)
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))

}

func main() {

}
