package model

import (
	"api-koperasi-backend/user/config"
	"api-koperasi-backend/user/shared"
	"crypto/sha256"
	
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/text/date"
)

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// ADD USER REGISTRATION
func CreateUser(email, password string) error {

	query := "INSERT INTO users ( email, password, is_verified) VALUES ( ?, ?, ?)"
	_, err := config.DB.Exec(query, email, password, false)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	// Menambah data ke tabel users_details
	// Menambah data ke tabel users_details
	queryDetails := `
	INSERT INTO users_details (
		nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, 
		status, agama, pekerjaan, alamat, kota, kode_pos, no_telepon,  
		verified, tanggal_bergabung, email, password, ktp, kk, rekening_koran, 
		prove_of_residence, jabatan
	) VALUES ('', '', '', '', 
	 '', '', '', '', '', '', '', 
	 'not_verified',NOW(), ?, ?, '', '', '', '', 'anggota')
	`

	// Menjalankan query untuk menambah data ke users_details
	_, err = config.DB.Exec(queryDetails, email, password)
	if err != nil {
	log.Println("Error inserting user details:", err)
	return err
	}


	return nil
}

// Case tambahan, jika user sudah login dan melakukan registrasi ulang dengan email/username yang sama
// start
func UpdateUserByEmail(email, newPassword string) error {
	query := `
        UPDATE users
        password = ?, is_verified = FALSE
        WHERE email = ? AND is_verified = FALSE`

	_, err := config.DB.Exec(query, newPassword, email)
	return err
}

func GetUserByEmail(email string) (*shared.RegisterRequest, error) {
	var user shared.RegisterRequest
	err := config.DB.QueryRow(`
        SELECT id, email, is_verified 
        FROM users 
        WHERE email = ?`, email).Scan(&user.ID, &user.Email, &user.IsVerified)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Println("No user found with email:", email)
			return nil, nil // Tidak ada pengguna dengan email tersebut
		}
		log.Println("Error executing query:", err)
		return nil, err
	}

	log.Printf("User found: %+v\n", user)
	return &user, nil
}

//end

func SaveVerificationToken(email, token string) error {
	// expired_at := time.Now().Add(24 * time.Hour) // 24 jam dari sekarang
	//expired_at := time.Now().Add(2 *time.Minute) //2 menit
	//location, _ := time.LoadLocation("Asia/Jakarta")
	//expired_at2 := time.Now().In(location).Add(2 * time.Minute) // Waktu dalam WIB
	expired_at := time.Now().UTC().Add(2 * time.Minute) // Waktu UTC

	log.Println("Expired At:", expired_at) // Log untuk memeriksa nilai
	// log.Println("kalo pake location jadinya:", expired_at2)
	// log.Println("kalo pake utc jadinya: ", expired_at3)

	query := "INSERT INTO verification_tokens (email, token, expired_at) VALUES (?, ?, ?)"
	_, err := config.DB.Exec(query, email, token, expired_at)

	log.Println("Saving token for email:", email, "Token:", token)

	if err != nil {
		log.Println("Error saving verification token:", err)
		return err
	}
	return nil
}

func VerifyUserEmail(token string) error {
	var email string
	var expired_at string

	// Ambil data dari database
	err := config.DB.QueryRow(`
    SELECT email, expired_at 
    FROM verification_tokens 
    WHERE token = ?`, token).Scan(&email, &expired_at)

	fmt.Println("tokennya : ", token)
	fmt.Println("usernya: ", email)
	fmt.Println("expired-nya: ", expired_at)

	if err != nil {
		fmt.Println("nilll, alasannya:", err)
		fmt.Println("tokennya : ", token)
		fmt.Println("usernya: ", email)
		fmt.Println("expired-nya: ", expired_at)

		if err.Error() == "sql:no rows in result set" {
			fmt.Println("tokennya gavalid")
			return errors.New("token tidak valid")
		}
		return err
	}

	log.Println("Verifying token:", token)
	log.Println("Found email:", email, "Token expiration:", expired_at)

	// Coba konversi expired_at dari string ke time.Time
	expiredAtParsed, err := time.Parse("2006-01-02 15:04:05", expired_at)
	if err != nil {
		return fmt.Errorf("gagal mengonversi expired_at ke time.Time: %v", err)
	}

	fmt.Println("Expired parsed:", expiredAtParsed)

	// Periksa apakah token sudah kadaluarsa
	if time.Now().After(expiredAtParsed) {
		return errors.New("token sudah kadaluarsa")
	}

	// Update status verifikasi pengguna menjadi true
	_, err = config.DB.Exec("UPDATE users SET is_verified = TRUE WHERE email = ?", email)
	if err != nil {
		return err
	}

	// Hapus token setelah berhasil digunakan
	_, err = config.DB.Exec("DELETE FROM verification_tokens WHERE token = ?", token)
	return err
}

// KEBUTUHAN LOGIN
// ValidateUser checks if the username and password are correct
func ValidateUser(username, password, jabatan string) (*shared.RegisterRequest, error) {
	var user shared.RegisterRequest

	// string Jabatan;
	err := config.DB.QueryRow(`
        SELECT id, email, password, jabatan
        FROM users 
        WHERE email = ? AND is_verified = TRUE`, username).
        Scan(&user.ID, &user.Email, &user.Password, &user.Jabatan)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username atau password salah atau akun belum diverifikasi")

		}
		log.Println("Error querying user:", err)
		return nil, err
	}

	// Periksa password menggunakan bcrypt
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("username atau password salah")
	}

	return &user, nil
}

func VerifyResetToken(token string, email *string) error {
	var expiredAt string
	err := config.DB.QueryRow(`
        SELECT email, expired_at 
        FROM verification_tokens 
        WHERE token = ?`, token).Scan(email, &expiredAt)

	if err != nil {
		return errors.New("token tidak valid")
	}

	// Periksa apakah token sudah kadaluarsa
	expiredAtParsed, err := time.Parse("2006-01-02 15:04:05", expiredAt)
	if err != nil {
		return fmt.Errorf("gagal mengonversi expired_at ke time.Time: %v", err)
	}

	if time.Now().After(expiredAtParsed) {
		return errors.New("token sudah kadaluarsa")
	}

	return nil
}

func UpdatePassword(email, hashedPassword string) error {
	_, err := config.DB.Exec("UPDATE users SET password = ? WHERE email = ?", hashedPassword, email)
	return err
}




//ACTION TAMBAH/REGISTRASI USER KE 2
func UpdateAnggota(user shared.Registrasi2) error {

	log.Println(user)
    query := `UPDATE users_details SET 
        nama_lengkap = ?, jenis_kelamin = ?, tempat = ?, tanggal_lahir = ?, 
        status = ?, agama = ?, pekerjaan = ?, alamat = ?, kota = ?, 
        kode_pos = ?, no_telepon = ?, verified = ?,
        ktp = ?, kk = ?, rekening_koran = ?, prove_of_residence = ?
        WHERE email = ?`

    _, err := config.DB.Exec(query,
        user.Nama_lengkap,
        user.Jenis_kelamin,
        user.Tempat,
        user.Tanggal_lahir,
        user.Status,
        user.Agama,
        user.Pekerjaan,
        user.Alamat,
        user.Kota,
        user.Kode_pos,
        user.No_telepon,
        user.Verified,
        user.Ktp,
        user.Kk,
        user.Rekening_koran,
        user.Prove_of_residence,
        user.Email, // Identifikasi berdasarkan email
    )
    if err != nil {
        log.Println("Error updating anggota:", err)
        return err
    }
    log.Println("Anggota data updated successfully")
    return nil
}
