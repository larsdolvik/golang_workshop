package challenge

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Post struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Solution fetches comments from the jsonplaceholder api and returns it to the user
func Solution(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var comments []Post

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1/comments")

	if err != nil {
		log.Printf("Reading response body failed. Reason: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	err = convertToStruct(&comments, resp)
	if err != nil {
		log.Printf("Converting to struct failed, reason: %v", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// convertToStruct takes a target interface and a http response and unmarshals the response to JSON
func convertToStruct(target interface{}, response *http.Response) error {
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Reading response body failed. Reason: %s", err.Error())
		return err
	}

	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		log.Printf("Unmarshaling response failed. Reason: %s", err.Error())
		return err
	}
	return nil
}
