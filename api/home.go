package api

import (
	"net/http"

	"github.com/imthaghost/scdl/cmd"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	cmd.Execute()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "HOME ROUTE"}`))

}
