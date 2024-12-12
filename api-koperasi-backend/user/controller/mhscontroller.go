package controller

import (
	// "encoding/json"
	// // "log"
	// "api-koperasi-backend/user/model"
	// // "myapp/shared"
	// "net/http"
	// "message"
	// "encoding/json"
	// "log"
	// "net/http"
	// "myapp/model"

)

type APIResponseMhs struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// func SaveStudent(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var student model.Student
// 	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
// 		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
// 		return
// 	}

// 	if err := model.CreateStudent(student); err != nil {
// 		http.Error(w, `{"error": "Failed to save student data"}`, http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "Student data saved successfully"}`))
// }