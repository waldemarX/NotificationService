package main

import (
	"notif/app/db"
	"notif/app/routes"
	"log"
	"net/http"
)


func main() {
	db.ConnectDB()
	http.HandleFunc("/like", routes.ClickLike)

	server_error := http.ListenAndServe(":8000", nil)
	if server_error != nil {
		log.Fatal("ListenAndServe: ", server_error)
	}
	log.Println("http server started on :8000")
}
