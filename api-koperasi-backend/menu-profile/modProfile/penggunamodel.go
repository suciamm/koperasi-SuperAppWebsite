package modprofile

import (
	"api-koperasi-backend/menu-profile/shared"
	"log"
	"api-koperasi-backend/menu-profile/confProfile"
	// "fmt"
	// "database/sql"
	// "errors"
	// "time"
    // "net/http"
    // "encoding/json"
)



// CreateStudent inserts a student into the database
// func CreateAnggota(user shared.Registrasi2) error {
// 	query := `INSERT INTO users_details (
//         nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, status, agama, pekerjaan, 
//         alamat, kota, kode_pos, no_telepon, verified, tanggal_bergabung, email, password, 
//         ktp, kk, rekening_koran, prove_of_residence, jabatan
//     ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

// // Eksekusi query
// _, err := config.DB.Exec(query,
// 	user.Nama_lengkap,
// 	user.Jenis_kelamin,
// 	user.Tempat,
// 	user.Tanggal_lahir,
// 	user.Status,
// 	user.Agama,
// 	user.Pekerjaan,
// 	user.Alamat,
// 	user.Kota,
// 	user.Kode_pos,
// 	user.No_telepon,
// 	user.Verified,
// 	user.Tanggal_bergabung,
// 	user.Email,
// 	user.Password,
// 	user.Ktp,
// 	user.Kk,
// 	user.Rekening_koran,
// 	user.Prove_of_residence,
// 	"", // Jabatan kosong sementara
// )

// 	if err != nil {
// 		log.Println("Error inserting anggota:", err)
// 		return err
// 	}
// 	return nil
// }


func CreateAnggota(user shared.Registrasi2) error {
    query := `INSERT INTO users_details (
        nama_lengkap, jenis_kelamin, tempat, tanggal_lahir, status, agama, pekerjaan, 
        alamat, kota, kode_pos, no_telepon, verified, tanggal_bergabung, email, password, 
        ktp, kk, rekening_koran, prove_of_residence, jabatan
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

    _, err := confprofile.DB.Exec(query,
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
        user.Tanggal_bergabung,
        user.Email,
        user.Password,
        user.Ktp,
        user.Kk,
        user.Rekening_koran,
        user.Prove_of_residence,
        "", // Jabatan kosong sementara
    )
    if err != nil {
        log.Println("Error inserting anggota:", err)
        return err
    }
    log.Println("Anggota data inserted successfully")
    return nil
}


// func GetUserProfileByEmail(email string) (*shared.Registrasi2, error) {
//     var admin shared.Registrasi2
//     var tanggalLahirStr, tanggalBergabungStr string
//     query := `
//         SELECT 
//             ud.id, ud.nama_lengkap, ud.jenis_kelamin, ud.tempat, ud.tanggal_lahir, ud.status, ud.agama, 
//             ud.pekerjaan, ud.alamat, ud.kota, ud.kode_pos, ud.no_telepon, 
//             u.is_verified AS verified, u.created_at AS tanggal_bergabung, 
//             u.email, u.password, 
//             ud.ktp, ud.kk, ud.rekening_koran, ud.prove_of_residence
//         FROM users_details ud
//         INNER JOIN users u ON ud.email = u.email
//         WHERE u.email = ? 
//     `
//     err := config.DB.QueryRow(query, email).Scan(
//         &admin.ID, &admin.Nama_lengkap, &admin.Jenis_kelamin, &admin.Tempat, &tanggalLahirStr,
//         &admin.Status, &admin.Agama, &admin.Pekerjaan, &admin.Alamat, &admin.Kota,
//         &admin.Kode_pos, &admin.No_telepon, &admin.Verified, &tanggalBergabungStr,
//         &admin.Email, &admin.Password, &admin.Ktp, &admin.Kk, &admin.Rekening_koran,
//         &admin.Prove_of_residence,
//     )
//     if err != nil {
//         if err == sql.ErrNoRows {
//             return nil, errors.New("data admin tidak ditemukan")
//         }
//         log.Println("Error querying admin:", err)
//         return nil, err
//     }
//     admin.Tanggal_lahir, err = time.Parse("2006-01-02", tanggalLahirStr)
//     if err != nil {
//         log.Printf("Error parsing tanggal_lahir: %v\n", err)
//         return nil, errors.New("gagal memproses tanggal lahir")
//     }
//     admin.Tanggal_bergabung, err = time.Parse("2006-01-02 15:04:05", tanggalBergabungStr)
//     if err != nil {
//         log.Printf("Error parsing tanggal_bergabung: %v\n", err)
//         return nil, errors.New("gagal memproses tanggal bergabung")
//     }
//     return &admin, nil
// }

