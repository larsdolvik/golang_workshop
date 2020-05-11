package hello

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SayHello(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}
