package user

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
)

var db *sql.DB

// InitDatabase melakukan koneksi ke database MySQL
func InitDatabase() {
	// Menghubungkan ke database lokal dengan user root dan password 'password'
	dsn := "root:root@tcp(localhost:8889)/koperasi_user_manajemen"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database:", err)
	}
	fmt.Println("database connect")
	// Verifikasi koneksi
	if err = db.Ping(); err != nil {
		log.Fatal("Koneksi gagal:", err)
	}
	fmt.Println("Koneksi ke database berhasil!")
}

// GetUserByEmail mengambil data pengguna berdasarkan email
func GetUserByEmail(email string) (map[string]interface{}, error) {
	query := `
		SELECT id, email, created_at, is_verified
		FROM users
		WHERE email = ?`
	
	var userId int
	var userEmail, createdAt string
	var isVerified bool

	err := db.QueryRow(query, email).Scan(&userId, &userEmail, &createdAt, &isVerified)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Pengguna tidak ditemukan
		}
		return nil, err
	}

	// Menyusun data dalam format map
	userData := map[string]interface{}{
		"id":          userId,
		"email":       userEmail,
		"created_at":  createdAt,
		"is_verified": isVerified,
	}
	return userData, nil
}

// GetVerificationToken mengembalikan token verifikasi berdasarkan email
func GetVerificationToken(email string) (string, error) {
	query := `
		SELECT token
		FROM verification_tokens
		WHERE email = ? AND expired_at > NOW()`
	
	var token string
	err := db.QueryRow(query, email).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // Token tidak ditemukan atau sudah expired
		}
		return "", err
	}

	return token, nil
}

// GetPasswordResetToken mengambil token reset password berdasarkan email
func GetPasswordResetToken(email string) (string, error) {
	query := `
		SELECT token
		FROM password_reset_tokens
		WHERE email = ? AND expired_at > NOW()`
	
	var token string
	err := db.QueryRow(query, email).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // Token tidak ditemukan atau sudah expired
		}
		return "", err
	}

	return token, nil
}
