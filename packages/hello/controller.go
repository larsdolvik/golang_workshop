package hello

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Hello returns a string saying Hello World
func Hello(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	/*
		The req and ps is not need in this example but it can be useful to see how
		they can be implemented as parameters.

		req is the request gotten from the client

		ps contains the urlparameters
	*/
	var helloworld string
	helloworld = "Hello World"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helloworld)
}

// SayHello returns a string saying Hello World
func SayHello(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	/*
		The req is not need in this example but it can be useful to see how
		they can be implemented as parameters.

		req is the request gotten from the client

		ps contains the urlparameters
	*/

	parameter := ps.ByName("name")

	var hello string
	hello = "Hello " + parameter

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}
