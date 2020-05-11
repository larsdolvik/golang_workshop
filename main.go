package main

import (
	"code/golang_workshop/packages/bye"
	"code/golang_workshop/packages/hello"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	router := httprouter.New()

	router.GET("/hello", hello.SayHello)
	router.POST("/bye", bye.SayBye)

	log.Println("Starting api...")

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", ""},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	http.ListenAndServe(":8080", handler)
}
