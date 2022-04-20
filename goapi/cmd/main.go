package main

import (
	"goapi/pkg/common/database"
	"goapi/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Test struct {
	Name string
}

func main() {
	// เรียกก่อนเริ่มเปิด server เพราะถ้าเชื่อมต่อไม่ได้ให้จะได้ไม่ต้อง start server
	database.ConnectDB()

	// เปลี่ยนตรงนี้
	r := mux.NewRouter()
	// define route
	setupRouter(r)

	// starting server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setupRouter(r *mux.Router) {
	todo := r.PathPrefix("/api/todos").Subrouter()
	todo.HandleFunc("", handlers.CreateTodo).Methods(http.MethodPost)
	todo.HandleFunc("", handlers.ListTodo).Methods(http.MethodGet)
	// สามารถใช้ร่วมกับ regx ได้
	todo.HandleFunc("/{id:[0-9]+}", handlers.GetTodo).Methods(http.MethodGet)
	todo.HandleFunc("/{id:[0-9]+}", handlers.UpdateTodoStatus).Methods(http.MethodPut)
	todo.HandleFunc("/{id:[0-9]+}", handlers.DeleteTodo).Methods(http.MethodDelete)
}
