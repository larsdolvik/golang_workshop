package bye

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SayBye(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Bye")
}
