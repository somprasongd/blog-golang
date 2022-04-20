package handlers

import (
	"encoding/json"
	"errors"
	"goapi/pkg/common/database"
	"goapi/pkg/common/validator"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text" gorm:"column:title" validate:"required" `
	Completed bool   `json:"isCompleted" gorm:"column:is_done"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// step 1: แปลง JSON จาก request body เป็น Todo struct
	var todo Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// step 2: validate
	err = validator.ValidateStruct(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// step 3: insert
	tx := database.DB.Create(&todo)
	if err := tx.Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// step 4: response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func ListTodo(w http.ResponseWriter, r *http.Request) {
	// เพิ่มอ่านค่าจาก query params
	query := r.URL.Query()
	wheres := map[string]interface{}{}
	// ถ้าส่ง completed มา ให้ใส่ไปใน map["is_done"] ตามชื่อ column จริงในฐานข้อมูล
	if val, ok := query["completed"]; ok {
		b1, err := strconv.ParseBool(val[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		wheres["is_done"] = b1
	}

	todos := []Todo{}
	// เพิ่ม .Where()
	tx := database.DB.Where(wheres).Find(&todos)

	if err := tx.Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	// step 1: get id from path param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// step 2: select where id
	todo := Todo{}
	tx := database.DB.First(&todo, id)
	if err := tx.Error; err != nil {
		// step 3: handle error not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "todo with given id not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// step 4: response
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodoStatus(w http.ResponseWriter, r *http.Request) {
	// step 1: get id from path param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// step 2: แปลง json body เป็น struct เพื่อเอาค่าสถานะที่ส่งมา
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// step 3: update only is_done column
	tx := database.DB.Model(Todo{ID: id}).Update("is_done", todo.Completed)
	if err := tx.Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// step 4: handle not found error
	if tx.RowsAffected == 0 {
		http.Error(w, "todo with given id not found", http.StatusNotFound)
		return
	}
	// step 5: response
	w.WriteHeader(http.StatusNoContent)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// step 1: get id from path param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// step 2: delete where id
	tx := database.DB.Delete(&Todo{}, id)
	if err := tx.Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// step 3: handle not found error
	if tx.RowsAffected <= 0 {
		http.Error(w, "todo with given id not found", http.StatusNotFound)
		return
	}
	// step 4: response
	w.WriteHeader(http.StatusNoContent)
}
