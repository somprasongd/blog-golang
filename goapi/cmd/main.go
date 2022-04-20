package main

import (
	"goapi/pkg/common/database"
	"goapi/pkg/handlers"
	"goapi/pkg/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	handler := setupRouter(r)

	// starting server
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func setupRouter(r *mux.Router) http.Handler {
	todo := r.PathPrefix("/api/todos").Subrouter()
	todoHandler := handlers.NewTodoHandler(database.DB)
	todo.HandleFunc("", todoHandler.CreateTodo).Methods(http.MethodPost)
	todo.HandleFunc("", todoHandler.ListTodo).Methods(http.MethodGet)
	// สามารถใช้ร่วมกับ regx ได้
	todo.HandleFunc("/{id:[0-9]+}", todoHandler.GetTodo).Methods(http.MethodGet)
	todo.HandleFunc("/{id:[0-9]+}", todoHandler.UpdateTodoStatus).Methods(http.MethodPut)
	todo.HandleFunc("/{id:[0-9]+}", todoHandler.DeleteTodo).Methods(http.MethodDelete)

	r.Use(middleware.Logging)

	// Handling CORS Requests
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	})

	handler := c.Handler(r)
	return handler
}
