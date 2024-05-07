package routes

import (
	"fmt"
	"log"
	"net/http"
)

func ClickLike(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Like added!")
	log.Printf("GET : %v", request.RequestURI)
}