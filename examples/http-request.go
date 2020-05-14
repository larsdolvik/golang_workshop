package examples

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Todo represents the data structure of a todo item
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// ShortHandRequest makes a http request todos
func ShortHandRequest() (Todo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	var result Todo

	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	err = convertToStruct(&result, resp)
	if err != nil {
		log.Printf("Converting to struct failed, reason: %v", err.Error())
	}

	return result, nil
}

// ToJSON takes a target interface and a http response and unmarshals the response to JSON
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
