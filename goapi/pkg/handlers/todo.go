package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"text"`
	Completed bool   `json:"completed" gorm:"column:is_done"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Todo")
}

func ListTodo(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{ID: 1, Title: "Test 1", Completed: true},
		{ID: 2, Title: "Test 2", Completed: false},
		{ID: 3, Title: "Test 3", Completed: false},
	}
	// ส่งข้อมูลทั้งหมดกลับไปในรูปแบบ JSON
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Get Todo by ID:", vars["id"])
}

func UpdateTodoStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Update Todo Status by ID:", vars["id"])
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Delete Todo by ID:", vars["id"])
}
