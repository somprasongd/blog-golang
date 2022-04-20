package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Test struct {
	Name string
}

func main() {
	// เปลี่ยนตรงนี้
	r := mux.NewRouter()
	// define route
	r.HandleFunc("/tests", handleTest)

	// starting server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(query)
}
