package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/pratapnarra/fetchapi/handlers"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()


	router.HandleFunc("/receipts/process",handlers.PostHandler).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetHandler).Methods("GET")


	log.Println("Server started on port 8080")
	fmt.Printf("Server started on port 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}