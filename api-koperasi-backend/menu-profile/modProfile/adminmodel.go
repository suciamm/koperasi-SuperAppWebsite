package modprofile

import (
	confprofile "api-koperasi-backend/menu-profile/confProfile"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
    "fmt"
	// "api-koperasi-backend/menu-profile/confProfile"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Struct untuk representasi data mahasiswa
type Anggota struct {
    ID       int    `json:"id"`
    Nama     string `json:"nama"`
    Email     string `json:"email"`
    Alamat  string `json:"alamat"`
    Verified string `json:"verified"`
    Tanggal_bergabung string `json:"tanggal_bergabung"`
}

type Users struct {
    Id               string       `json:"id"`
    Nama_lengkap    string    `json:"nama_lengkap"`
    Jenis_kelamin   string    `json:"jenis_kelamin"`
    Tempat          string    `json:"tempat"`
    Tanggal_lahir   time.Time `json:"tanggal_lahir"`
    Status          string    `json:"status"`
    Agama           string    `json:"agama"`
    Pekerjaan       string    `json:"pekerjaan"`
    Alamat          string    `json:"alamat"`
    Kota            string    `json:"kota"`
    Kode_pos        string       `json:"kode_pos"`
    No_telepon      string    `json:"no_telepon"`
    Verified        string    `json:"verified"`
    Tanggal_bergabung time.Time `json:"tanggal_bergabung"`
    Email           string    `json:"email"`
    Password        string    `json:"password"`
    Ktp             string    `json:"ktp"`
    Kk              string    `json:"kk"`
    Rekening_koran  string    `json:"rekening_koran"`
    Prove_of_residence string `json:"prove_of_residence"`
}



// Fungsi untuk handle request GET /api/Anggota
func GetDataAnggota(w http.ResponseWriter, r *http.Request) {
    // Koneksi ke database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/koperasi_user_manajemen")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Query ke database
    rows, err := db.Query("SELECT nama_lengkap, alamat, verified FROM users_details")
    if err != nil {
        http.Error(w, "Gagal mengambil data dari database", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var anggota []Anggota
    for rows.Next() {
        var mhs Anggota
        if err := rows.Scan(&mhs.Nama, &mhs.Alamat, &mhs.Verified); err != nil {
            http.Error(w, "Gagal memproses data anggota", http.StatusInternalServerError)
            return
        }
        anggota = append(anggota, mhs)
    }

    // Mengubah data ke JSON dan mengirimkan sebagai response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(anggota)
}



// func GetDataAnggota(w http.ResponseWriter, r *http.Request) {
//     rows, err := config.DB.Query("SELECT nama_lengkap, alamat, verified FROM users_details")
//     if err != nil {
//         http.Error(w, "Gagal mengambil data dari database", http.StatusInternalServerError)
//         return
//     }
//     defer rows.Close()

//     var anggota []Anggota
//     for rows.Next() {
//         var mhs Anggota
//         if err := rows.Scan(&mhs.Nama, &mhs.Alamat, &mhs.Verified); err != nil {
//             http.Error(w, "Gagal memproses data anggota", http.StatusInternalServerError)
//             return
//         }
//         anggota = append(anggota, mhs)
//     }

//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(anggota)
// }


func GetDataDaftarAnggota(w http.ResponseWriter, r *http.Request) {
    // Koneksi ke database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/koperasi_user_manajemen")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Query ke database
    rows, err := db.Query("SELECT `id`,`nama_lengkap`, `email`, `alamat`,`verified` FROM users_details WHERE verified = 'not_verified'")
    if err != nil {
        http.Error(w, "Gagal mengambil data dari database", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var anggota []Anggota
    for rows.Next() {
        var mhs Anggota
        if err := rows.Scan(&mhs.ID, &mhs.Nama, &mhs.Email,&mhs.Alamat, &mhs.Verified); err != nil {
            http.Error(w, "Gagal memproses data anggota", http.StatusInternalServerError)
            return
        }
        anggota = append(anggota, mhs)
    }

    // Mengubah data ke JSON dan mengirimkan sebagai response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(anggota)
}





//detail pendaftaran
func GetUserDetails(w http.ResponseWriter, r *http.Request) {

    // Ambil ID dari path
    id := mux.Vars(r)["id"]

    // Query database untuk ID tertentu
    var anggota Users
    err := confprofile.DB.QueryRow("SELECT nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, status, agama, pekerjaan, alamat, kota, kode_pos, no_telepon, verified, tanggal_bergabung, email, ktp, kk, rekening_koran, prove_of_residence FROM users_details WHERE id = ?", id).
    Scan(&anggota.Nama_lengkap, &anggota.Jenis_kelamin, &anggota.Tempat, &anggota.Tanggal_lahir, &anggota.Status, &anggota.Agama, &anggota.Pekerjaan, &anggota.Alamat, &anggota.Kota, &anggota.Kode_pos, &anggota.No_telepon, &anggota.Verified, &anggota.Tanggal_bergabung, &anggota.Email, &anggota.Ktp, &anggota.Kk, &anggota.Rekening_koran, &anggota.Prove_of_residence)


    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Data anggota tidak ditemukan", http.StatusNotFound)
        } else {
            http.Error(w, "Gagal mengambil data anggota", http.StatusInternalServerError)
        }
        return
    }
}


//DEATAIL PENDAFTARAN ANGGOTA
// Fungsi handler untuk mendapatkan data anggota berdasarkan ID
func GetAnggotaByID(w http.ResponseWriter, r *http.Request) {
	// Parsing parameter ID dari URL (query string)
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "ID anggota tidak ditemukan", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	// Koneksi ke database
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/koperasi_user_manajemen")
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}
	defer db.Close()

	// Query untuk mendapatkan data anggota berdasarkan ID
	query := `SELECT id, nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, status, agama, pekerjaan, 
	alamat, kota, kode_pos, no_telepon, verified, tanggal_bergabung, email, ktp, kk, rekening_koran, prove_of_residence 
	FROM users_details WHERE id = ?`

	row := db.QueryRow(query, id)

	// Memindahkan hasil query ke dalam struct
	var user Users
	err = row.Scan(
		&user.Id, &user.Nama_lengkap, &user.Jenis_kelamin, &user.Tempat, &user.Tanggal_lahir, &user.Status,
		&user.Agama, &user.Pekerjaan, &user.Alamat, &user.Kota, &user.Kode_pos, &user.No_telepon,
		&user.Verified, &user.Tanggal_bergabung, &user.Email, &user.Ktp, &user.Kk, &user.Rekening_koran, &user.Prove_of_residence,
	)

	// Jika tidak ada data yang ditemukan 
	if err == sql.ErrNoRows {
		http.Error(w, "Anggota tidak ditemukan", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Gagal mendapatkan data", http.StatusInternalServerError)
		return
	}
	// Mengatur header untuk response JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Mengembalikan data dalam format JSON
	json.NewEncoder(w).Encode(user)
}
 


//GET DATA ICON
// Fungsi untuk handle request GET /api/Anggota
func GetDataIcon(w http.ResponseWriter, r *http.Request) {
    // Koneksi ke database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/koperasi_user_manajemen")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Query ke database
    rows, err := db.Query("SELECT nama_lengkap, alamat, verified FROM users_details")
    if err != nil {
        http.Error(w, "Gagal mengambil data dari database", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var anggota []Anggota
    for rows.Next() {
        var mhs Anggota
        if err := rows.Scan(&mhs.Nama, &mhs.Alamat, &mhs.Verified); err != nil {
            http.Error(w, "Gagal memproses data anggota", http.StatusInternalServerError)
            return
        }
        anggota = append(anggota, mhs)
    }

    // Mengubah data ke JSON dan mengirimkan sebagai response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(anggota)
}



//Tampilkan detail anggota di daftar pengjuan
func GetDataDaftarAnggotaDetail(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/koperasi_user_manajemen")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    id := r.URL.Query().Get("id")
    fmt.Printf("ID yang diterima: %s\n", id)
    if id == "" {
        http.Error(w, "Missing id parameter", http.StatusBadRequest)
        return
    }
    // Ambil parameter email dari query string
	// id := r.URL.Query().Get("id")
    log.Println("idnya adalahhhh --------------------: ", id)
    log.Println("idnya adalahhhh --------------------: ", id)
    log.Println("idnya adalahhhh --------------------: ", id)
    log.Println("idnya adalahhhh --------------------: ", id)
    log.Println("idnya adalahhhh --------------------: ", id)
    log.Println("idnya adalahhhh --------------------: ", id)

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
    query := `
		SELECT id, nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, 
		       status, agama, pekerjaan, alamat, kota, kode_pos, no_telepon, 
		       email, verified, tanggal_bergabung, ktp, kk, rekening_koran, prove_of_residence
		FROM users_details 
		WHERE id = ?
	`
	
	// Eksekusi query
	row := db.QueryRow(query, id)

	// Struct untuk menampung hasil query
	var user struct {
        ID               string `json:"id"`
		NamaLengkap      string `json:"nama_lengkap"`
		JenisKelamin     string `json:"jenis_kelamin"`
		Tempat           string `json:"tempat"`
		TanggalLahir     string `json:"tanggal_lahir"`
		Status           string `json:"status"`
		Agama            string `json:"agama"`
		Pekerjaan        string `json:"pekerjaan"`
		Alamat           string `json:"alamat"`
		Kota             string `json:"kota"`
		KodePos          string `json:"kode_pos"`
		NoTelepon        string `json:"no_telepon"`
		Email            string `json:"email"`
		Verified         string `json:"verified"`
		TanggalBergabung string `json:"tanggal_bergabung"`
		KTP              string `json:"ktp"`
		KK               string `json:"kk"`
		RekeningKoran    string `json:"rekening_koran"`
		ProveOfResidence string `json:"prove_of_residence"`
        
	}

	// Scan hasil query ke struct
	err = row.Scan(
		&user.ID, &user.NamaLengkap, &user.JenisKelamin, &user.Tempat, &user.TanggalLahir,
		&user.Status, &user.Agama, &user.Pekerjaan, &user.Alamat, &user.Kota, &user.KodePos,
		&user.NoTelepon, &user.Email, &user.Verified, &user.TanggalBergabung,
		&user.KTP, &user.KK, &user.RekeningKoran, &user.ProveOfResidence,
	)
    log.Println("----------------------")
    log.Println("----------------------")
    log.Println("----------------------")
    log.Println("----------------------")
    log.Println(row.Scan(user.NamaLengkap))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kirim data sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
//INSERT DATA ANGGOTA




