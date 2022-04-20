package main

import (
	"encoding/json"
	"fmt"
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
	// สามารถใช้ร่วมกับ regx ได้
	r.HandleFunc("/tests/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprint(w, vars["id"])
	})

	// starting server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(query)
}
