package controller

import (
	"encoding/json"
	"database/sql"
	"api-koperasi-backend/user/config"
	"api-koperasi-backend/user/model"
	"api-koperasi-backend/user/shared"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"regexp"
	"io"
	"golang.org/x/crypto/bcrypt"
)

// // // Struct untuk menerima data JSON
// type RegisterRequest struct {
// 	ID 				int
// 	Email           string `json:"email"`
// 	Password        string `json:"password"`
// 	ConfirmPassword string `json:"confirm_password"`

// 	IsVerified bool
//     Jabatan string

// }

// Middleware untuk menangani CORS
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Izinkan semua origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Jika metode OPTIONS, cukup kembalikan header tanpa meneruskan ke handler berikutnya
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Struct untuk mengirim respon JSON
type APIResponse struct {
	Error string `json:"error,omitempty"`
}

// HashPassword meng-hash password menggunakan bcrypt
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}
	return string(hashedPassword)
}

// function untuk validasi password harus 6 karakter
func validatePassword(password string) bool {
	// Password harus terdiri dari minimal 6 karakter
	if len(password) < 6 {
		return false
	}

	// Memeriksa apakah password mengandung huruf besar, huruf kecil, dan karakter khusus
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString // Menggunakan \W untuk karakter non-alphanumeric (termasuk spasi dan simbol)

	return hasUpper(password) && hasLower(password) && hasSpecial(password)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("RegisterUser handler called")

	if r.Method == http.MethodPost {
		// Parse body JSON
		var req shared.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validasi password dan konfirmasi password
		if req.Password != req.ConfirmPassword {
			http.Error(w, `{"error": "Passwords dan konfirmasi password tidak sama"}`, http.StatusBadRequest)
			return
		}

		// Validasi panjang password
		if !validatePassword(req.Password) {
			http.Error(w, `{"error": "Passwords harus terdiri dari 6 karakter, mengandung huruf besar, huruf kecil, dan karakter khusus."}`, http.StatusBadRequest)
			return
		}

		// Validasi cek email sudah terdaftar di database?
		existingUser, err := model.GetUserByEmail(req.Email)
		if err != nil {
			log.Println("Error checking user:", err)
			http.Error(w, `{"error": "Gagal memeriksa data pengguna"}`, http.StatusBadRequest)
			return
		}

		// Validasi user terlambat klik link verifikasi email
		if existingUser != nil && !existingUser.IsVerified {
			// Update data pengguna lama
			err := model.UpdateUserByEmail(req.Email, HashPassword(req.Password))
			if err != nil {
				log.Println("Error updating user:", err)
				// pageVariables := RegisterPageData{
				http.Error(w, `{"error": "Gagal memperbarui data pengguna lama"}`, http.StatusBadRequest)
				// }
				// renderRegisterPage(w, pageVariables)
				return
			}
		} else if existingUser == nil {
			// Buat pengguna baru
			err := model.CreateUser(req.Email, HashPassword(req.Password))
			if err != nil {
				log.Println("Error creating user:", err)
				// pageVariables := RegisterPageData{
				http.Error(w, `{"error": "Gagal menyimpan data pengguna"}`, http.StatusBadRequest)
				// }
				// renderRegisterPage(w, pageVariables)
				return
			}
		} else {
			// pageVariables := RegisterPageData{
			http.Error(w, `{"error": "Email sudah digunakan dan terverifikasi"}`, http.StatusBadRequest)
			// }
			// renderRegisterPage(w, pageVariables)
			return
		}

		// Kirim email verifikasi
		verificationToken := generateVerificationToken()
		err = sendVerificationEmail(req.Email, verificationToken)
		if err != nil {
			// pageVariables := RegisterPageData{
			http.Error(w, `{"error": "Gagal mengirim email verifikasi"}`, http.StatusBadRequest)
			// }
			// renderRegisterPage(w, pageVariables)
			return
		}

		// Simpan token verifikasi di database
		err = model.SaveVerificationToken(req.Email, verificationToken)
		if err != nil {
			// pageVariables := RegisterPageData{
			http.Error(w, `{"error": "Gagal menyimpan token verifikasi"}`, http.StatusBadRequest)
			// }
			// renderRegisterPage(w, pageVariables)
			return
		}

		// Jika berhasil, kirim respon sukses
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(APIResponse{})
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// GENERATE TOKEN EMAIL
func generateVerificationToken() string {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		fmt.Println("berhasil generate")
		log.Fatal("Error generating token:", err)
	}
	return hex.EncodeToString(token)
}

// SENDING EMAIL VERIFICATION
func sendVerificationEmail(to, token string) error {
	from := "systexsuci22@gmail.com"
	password := "lqocianpnkykmwqq"
	smtpServer := "smtp.gmail.com"
	port := "587"
	// port := "465"
	subject := "Email Verifikasi"
	body := fmt.Sprintf("Klik link berikut untuk memverifikasi email Anda: http://localhost:8080/verify?token=%s", token)
	log.Println("Received token:", token)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Log untuk membantu debugging
	log.Println("Sending email to:", to)
	log.Println("SMTP server:", smtpServer+":"+port)

	err := smtp.SendMail(smtpServer+":"+port, auth, from, []string{to}, msg)
	if err != nil {
		log.Println("Error sending email:", err)
	}

	return smtp.SendMail(smtpServer+":"+port, auth, from, []string{to}, msg)
}

// TIMEOUT VERIFIKASI LINKING EMAIL REGISTERED
func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token tidak ditemukan", http.StatusBadRequest)
		return
	}

	err := model.VerifyUserEmail(token)
	if err != nil {
		if err.Error() == "token sudah kedaluwarsa" {
			http.Error(w, "Token sudah kedaluwarsa. Silakan mendaftar ulang.", http.StatusUnauthorized)
			return
		}
		log.Println("Token received for verification:", token)
		http.Error(w, "Verifikasi gagal", http.StatusInternalServerError)
		return
	}

	// Redirect ke halaman login dengan pesan sukses
	http.Redirect(w, r, "/login?message=Email%20Anda%20berhasil%20diverifikasi.%20Silakan%20login.", http.StatusSeeOther)

	// fmt.Fprintf(w, "Email Anda berhasil diverifikasi. Silakan login.")
}

// FUNCTION UNTUK RESET PASSWORD
func RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")

		// Periksa apakah email ada di database
		user, err := model.GetUserByEmail(email)
		if err != nil {
			http.Error(w, "Gagal memeriksa data pengguna", http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "Email tidak terdaftar", http.StatusBadRequest)
			return
		}

		// Buat token reset password
		resetToken := generateVerificationToken()

		// Simpan token reset password
		err = model.SaveVerificationToken(email, resetToken)
		if err != nil {
			http.Error(w, "Gagal menyimpan token reset password", http.StatusInternalServerError)
			return
		}

		// Kirim email reset password
		resetLink := fmt.Sprintf("http://localhost:8010/form-reset?token=%s", resetToken)
		err = sendResetPasswordEmail(email, resetLink)
		if err != nil {
			http.Error(w, "Gagal mengirim email reset password", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Link reset password telah dikirim ke email Anda.")
	}
}

func sendResetPasswordEmail(to, resetLink string) error {
	from := "systexsuci22@gmail.com"
	password := "lqocianpnkykmwqq"
	smtpServer := "smtp.gmail.com"
	port := "587"
	subject := "Reset Password"
	body := fmt.Sprintf("Klik link berikut untuk mereset password Anda: %s", resetLink)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpServer)
	return smtp.SendMail(smtpServer+":"+port, auth, from, []string{to}, msg)
}

func ShowResetPasswordForm(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token tidak valid", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("views/form-reset.html")
	if err != nil {
		http.Error(w, "Gagal memuat template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
		"Token": token,
	})
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		token := r.FormValue("token")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		if password != confirmPassword {
			http.Error(w, "Password dan konfirmasi password tidak cocok", http.StatusBadRequest)
			return
		}

		// Periksa token dan dapatkan email
		var email string
		err := model.VerifyResetToken(token, &email)
		if err != nil {
			http.Error(w, "Token tidak valid atau kadaluarsa", http.StatusBadRequest)
			return
		}

		// Hash password baru
		hashedPassword := HashPassword(password)

		// Perbarui password pengguna
		err = model.UpdatePassword(email, hashedPassword)
		if err != nil {
			http.Error(w, "Gagal memperbarui password", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Password berhasil diperbarui!")
	}
}

//------------
//Function action untuk menangani login
//----------------------

type LoginPageData struct {
	Error string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Render the login page. Initially, no error is passed.
	pageVariables := LoginPageData{}
	// gbr := "base64xxx"
	tmpl, err := template.ParseFiles("views/login.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, pageVariables)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var loginRequest shared.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			fmt.Printf("error: %s\n", fmt.Sprintf("gagal decode request: %v", err))
			return
		}

		email := loginRequest.Email
		password := loginRequest.Password
		jabatan := loginRequest.Jabatan

		fmt.Println()
		user, err := model.ValidateUser(email, password, jabatan)
		if err != nil {
			log.Println("Login error:", err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Email atau password salah"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Login berhasil",
			"email": user.Email,
			"password": user.Password,
			"jabatan": user.Jabatan,
		})
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}











func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter email dari query string
	email := r.URL.Query().Get("email")
    log.Println("emailnya adalahhhh --------------------: ", email)
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
    query := `
		SELECT nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, 
		       status, agama, pekerjaan, alamat, kota, kode_pos, no_telepon, 
		       email, verified, tanggal_bergabung, ktp, kk, rekening_koran, prove_of_residence
		FROM users_details 
		WHERE email = ?
	`
	
	// Eksekusi query
	row := config.DB.QueryRow(query, email)

	// Struct untuk menampung hasil query
	var user struct {
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
	err := row.Scan(
		&user.NamaLengkap, &user.JenisKelamin, &user.Tempat, &user.TanggalLahir,
		&user.Status, &user.Agama, &user.Pekerjaan, &user.Alamat, &user.Kota, &user.KodePos,
		&user.NoTelepon, &user.Email, &user.Verified, &user.TanggalBergabung,
		&user.KTP, &user.KK, &user.RekeningKoran, &user.ProveOfResidence,
	)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kirim data sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}







func GetUserProfileDetail(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter ID dari query string
	id := r.URL.Query().Get("id")
	log.Println("ID yang diterima adalah --------------------: ", id)

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Query untuk mendapatkan detail user berdasarkan ID
	query := `
		SELECT nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, 
		       status, agama, pekerjaan, alamat, kota, kode_pos, no_telepon, 
		       email, verified, tanggal_bergabung, ktp, kk, rekening_koran, prove_of_residence
		FROM users_details 
		WHERE id = ?
	`

	// Eksekusi query
	row := config.DB.QueryRow(query, id)

	// Struct untuk menampung hasil query
	var user struct {
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
	err := row.Scan(
		&user.NamaLengkap, &user.JenisKelamin, &user.Tempat, &user.TanggalLahir,
		&user.Status, &user.Agama, &user.Pekerjaan, &user.Alamat, &user.Kota, &user.KodePos,
		&user.NoTelepon, &user.Email, &user.Verified, &user.TanggalBergabung,
		&user.KTP, &user.KK, &user.RekeningKoran, &user.ProveOfResidence,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Println("Error scanning data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Kirim data sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}



// func RegisterUser2(w http.ResponseWriter, r *http.Request) {
// 	//cek data yang masuk:
// 	body, _ := io.ReadAll(r.Body)
// 	log.Println("Request body:", string(body))



//     if r.Method != http.MethodPut { // Ganti POST dengan PUT
//         http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
//         return
//     }

// 	log.Println("--------------- : berhasil masuk ke function controller")
// 	// log.Println(r.Body)
//     var anggota shared.Registrasi2
//     if err := json.NewDecoder(r.Body).Decode(&anggota); err != nil {
//         log.Println("Error decoding JSON:", err)
// 		log.Println(r.Body)
//         http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
//         return
//     }

//     log.Printf("Data diterima untuk update: %+v\n", anggota)

//     // Update data di database
//     if err := model.UpdateAnggota(anggota); err != nil {
//         http.Error(w, `{"error": "Failed to update anggota data"}`, http.StatusInternalServerError)
//         return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     response := map[string]string{"message": "Data berhasil diperbarui"}
//     json.NewEncoder(w).Encode(response)
// }

func RegisterUser2(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }

    log.Println("--------------- : berhasil masuk ke function controller")

    // Parse multipart form
    if err := r.ParseMultipartForm(10 << 20); err != nil { // Maks 10 MB
        log.Println("Error parsing form:", err)
        http.Error(w, `{"error": "Invalid form data"}`, http.StatusBadRequest)
        return
    }

	//cek data yang masuk sebelum diubah
	body, _ := io.ReadAll(r.Body)
	log.Println("Raw request body:", string(body))


    anggota := shared.Registrasi2{
        Nama_lengkap:     r.FormValue("nama_lengkap"),
        Jenis_kelamin:    r.FormValue("jenis_kelamin"),
        Tempat:           r.FormValue("tempat"),
        Tanggal_lahir:    r.FormValue("tanggal_lahir"),
        Status:           r.FormValue("status"),
        Agama:            r.FormValue("agama"),
        Pekerjaan:        r.FormValue("pekerjaan"),
        Alamat:           r.FormValue("alamat"),
        Kota:             r.FormValue("kota"),
        Kode_pos:         r.FormValue("kode_pos"),
        No_telepon:       r.FormValue("no_telepon"),
        Email:            r.FormValue("email"),
        Verified:         r.FormValue("verified"),
        Ktp:              r.FormValue("ktp"),
        Kk:               r.FormValue("kk"),
        Rekening_koran:   r.FormValue("rekening_koran"),
        Prove_of_residence: r.FormValue("prove_of_residence"),
    }
    log.Printf("Data diterima untuk update: %+v\n", anggota)

	
	validValues := map[string]bool{"verified": true, "not_verified": true}
	if !validValues[anggota.Verified] {
		log.Println("Invalid value for verified:", anggota.Verified)
		anggota.Verified = "not_verified" // Default value jika tidak valid
	}


    if err := model.UpdateAnggota(anggota); err != nil {
        log.Println("Error updating anggota:", err)
        http.Error(w, `{"error": "Failed to update anggota data"}`, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Data berhasil diperbarui"})
}

