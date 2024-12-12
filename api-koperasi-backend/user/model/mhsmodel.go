package model

import (
	"database/sql"
	"encoding/json"
	"log"

	// "log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Struct untuk representasi data mahasiswa
type Mahasiswa struct {
	ID       string `json:"id"`
	Nama     string `json:"nama"`
	Jurusan  string `json:"jurusan"`
	Fakultas string `json:"fakultas"`
}

// Fungsi untuk koneksi ke database
func connectDB() (*sql.DB, error) {
	// Ganti root:root dan localhost:8889 sesuai konfigurasi MySQL Anda
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/universitas")
	if err != nil {
		return nil, err
	}
	// Tes koneksi
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Fungsi untuk mengambil semua data mahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Gagal terhubung ke database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nama, jurusan, fakultas FROM mahasiswa")
	if err != nil {
		http.Error(w, "Gagal mengambil data dari database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mahasiswas []Mahasiswa
	for rows.Next() {
		var mhs Mahasiswa
		if err := rows.Scan(&mhs.ID, &mhs.Nama, &mhs.Jurusan, &mhs.Fakultas); err != nil {
			http.Error(w, "Gagal memproses data mahasiswa", http.StatusInternalServerError)
			return
		}
		mahasiswas = append(mahasiswas, mhs)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mahasiswas)
}

// Fungsi untuk memperbarui data mahasiswa berdasarkan ID
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	log.Println("COMMINGGG DI CONTROLLER MAHASISWA UPDATE")
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Gagal terhubung ke database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var mhs Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		http.Error(w, "Input tidak valid", http.StatusBadRequest)
		return
	}

	query := "UPDATE mahasiswa SET nama = ?, jurusan = ?, fakultas = ? WHERE id = ?"
	result, err := db.Exec(query, mhs.Nama, mhs.Jurusan, mhs.Fakultas, mhs.ID)
	if err != nil {
		http.Error(w, "Gagal memperbarui data mahasiswa", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Mahasiswa tidak ditemukan", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data mahasiswa berhasil diperbarui"))
}
