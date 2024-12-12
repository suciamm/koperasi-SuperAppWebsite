package user

import (
	"encoding/json"
	"net/http"
)

// UserHandler menangani request untuk mendapatkan data pengguna berdasarkan email
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan parameter email dari query string
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email tidak ditemukan", http.StatusBadRequest)
		return
	}

	// Mengambil data pengguna berdasarkan email
	user, err := GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Gagal mengambil data pengguna", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "Pengguna tidak ditemukan", http.StatusNotFound)
		return
	}

	// Mengambil token verifikasi dan token reset password
	verificationToken, err := GetVerificationToken(email)
	if err != nil {
		http.Error(w, "Gagal mengambil token verifikasi", http.StatusInternalServerError)
		return
	}

	resetToken, err := GetPasswordResetToken(email)
	if err != nil {
		http.Error(w, "Gagal mengambil token reset password", http.StatusInternalServerError)
		return
	}

	// Menambahkan token ke data pengguna
	user["verification_token"] = verificationToken
	user["password_reset_token"] = resetToken

	// Mengembalikan data dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
