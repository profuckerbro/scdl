package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	rl "github.com/ahmedash95/ratelimit"
)

var ratelimit rl.Limit

func RunServ() {

	ratelimit = rl.CreateLimit("1r/s")

	r := mux.NewRouter().StrictSlash(true)
	r.Handle("/download",RateLimitMiddleware(http.HandlerFunc(downloadHandler)))
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))

}


//MIDDLEWARE
func RateLimitMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := "127.0.0.1" // use ip or user agent any key you want
		if !isValidRequest(ratelimit, ip) {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		ratelimit.Hit(ip)
		h.ServeHTTP(w, r)
	})
}

func isValidRequest(l rl.Limit, key string) bool {
	_, ok := l.Rates[key]
	if !ok {
		return true
	}
	if l.Rates[key].Hits == l.MaxRequests {
		return false
	}
	return true
}