package controllerProfile

import (
	"encoding/json"
	"log"
	"api-koperasi-backend/menu-profile/modProfile"
	// "myapp/shared"
	"api-koperasi-backend/menu-profile/shared"
	"net/http"
	// "template"
    // "api-koperasi-backend/menu-profile/confProfile"

)

type APIResponseMhs struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func SaveAnggota(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }

    var anggota shared.Registrasi2
    if err := json.NewDecoder(r.Body).Decode(&anggota); err != nil {
        log.Println("Error decoding JSON:", err)
        http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
        return
    }

    log.Printf("Data diterima: %+v\n", anggota)

    // Menyimpan data ke database
    if err := modprofile.CreateAnggota(anggota); err != nil {
        http.Error(w, `{"error": "Failed to save anggota data"}`, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    response := APIResponseMhs{Message: "Data berhasil disimpan"}
    json.NewEncoder(w).Encode(response)
}


// //panggil data pengguna
// func AdminProfileHandler(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "user-session")
//     email, ok := session.Values["email"].(string)

//     if !ok {
//         http.Redirect(w, r, "/login", http.StatusSeeOther)
//         return
//     }

//     admin, err := modprofile.GetSPAdminByEmail(email)
//     if err != nil {
//         log.Println("Error retrieving admin:", err)
//         http.Error(w, "Unable to retrieve admin data", http.StatusInternalServerError)
//         return
//     }

//     tmpl, err := template.ParseFiles("views/sp/admin/profile.html")
//     if err != nil {
//         log.Println("Error loading template:", err)
//         http.Error(w, "Unable to load template", http.StatusInternalServerError)
//         return
//     }

//     tmpl.Execute(w, admin)
// }





// func GetUserProfile(w http.ResponseWriter, r *http.Request) {
// 	// Ambil parameter email dari query string
// 	email := r.URL.Query().Get("email")
//     log.Println("emailnya adalahhhh --------------------: ", email)
// 	if email == "" {
// 		http.Error(w, "Email is required", http.StatusBadRequest)
// 		return
// 	}

// 	// Query untuk mengambil data user berdasarkan email
// 	// query := `
// 	// 	SELECT nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, 
// 	// 	       status, agama, pekerjaan, alamat, kota, kode_pos, no_telepon, 
// 	// 	       email, verified, tanggal_bergabung, ktp, kk, rekening_koran, prove_of_residence
// 	// 	FROM users_details 
// 	// 	WHERE email = ?
// 	// `
//     query := `
// 		SELECT nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, 
// 		       status, agama, pekerjaan, alamat, kota, kode_pos, no_telepon, 
// 		       email, verified, tanggal_bergabung, ktp, kk, rekening_koran, prove_of_residence
// 		FROM users_details 
// 		WHERE email = ?
// 	`

// 	// Eksekusi query
// 	row := confprofile.DB.QueryRow(query, email)

// 	// Struct untuk menampung hasil query
// 	var user struct {
// 		NamaLengkap      string `json:"nama_lengkap"`
// 		JenisKelamin     string `json:"jenis_kelamin"`
// 		Tempat           string `json:"tempat"`
// 		TanggalLahir     string `json:"tanggal_lahir"`
// 		Status           string `json:"status"`
// 		Agama            string `json:"agama"`
// 		Pekerjaan        string `json:"pekerjaan"`
// 		Alamat           string `json:"alamat"`
// 		Kota             string `json:"kota"`
// 		KodePos          string `json:"kode_pos"`
// 		NoTelepon        string `json:"no_telepon"`
// 		Email            string `json:"email"`
// 		Verified         string `json:"verified"`
// 		TanggalBergabung string `json:"tanggal_bergabung"`
// 		KTP              string `json:"ktp"`
// 		KK               string `json:"kk"`
// 		RekeningKoran    string `json:"rekening_koran"`
// 		ProveOfResidence string `json:"prove_of_residence"`
// 	}

// 	// Scan hasil query ke struct
// 	err := row.Scan(
// 		&user.NamaLengkap, &user.JenisKelamin, &user.Tempat, &user.TanggalLahir,
// 		&user.Status, &user.Agama, &user.Pekerjaan, &user.Alamat, &user.Kota, &user.KodePos,
// 		&user.NoTelepon, &user.Email, &user.Verified, &user.TanggalBergabung,
// 		&user.KTP, &user.KK, &user.RekeningKoran, &user.ProveOfResidence,
// 	)
// 	if err != nil {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	// Kirim data sebagai JSON
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }




