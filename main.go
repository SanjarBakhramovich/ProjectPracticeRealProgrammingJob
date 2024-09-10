package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message = "global variable"

type MessageRequest struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", message)
}

func UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req MessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	message = req.Message
	fmt.Fprintln(w, "Message updated successfully")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/update-message", UpdateMessageHandler).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}