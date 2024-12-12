package shared

import (
	"time"
)

type Registrasi2 struct {
	ID				   int		 `json:"id"`
	Nama_lengkap       string    `json:"nama_lengkap"`
	Jenis_kelamin      string    `json:"jenis_kelamin"`
	Tempat             string    `json:"tempat"`
	Tanggal_lahir      time.Time `json:"tanggal_lahir"`
	Status             string    `json:"status"`
	Agama              string    `json:"agama"`
	Pekerjaan          string    `json:"pekerjaan"`
	Alamat             string    `json:"alamat"`
	Kota               string    `json:"kota"`
	Kode_pos           int       `json:"kode_pos"`
	No_telepon         string    `json:"no_telepon"`
	Verified           string    `json:"verified"`
	Tanggal_bergabung  time.Time `json:"tanggal_bergabung"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Ktp                string    `json:"ktp"`
	Kk                 string    `json:"kk"`
	Rekening_koran     string    `json:"rekening_koran"`
	Prove_of_residence string    `json:"prove_of_residence"`
}  

// type Registrasi2 struct {
// 	Nama_lengkap       string    `json:"nama_lengkap"`
// 	Jenis_kelamin      string    `json:"jenis_kelamin"`
// 	Tempat             string    `json:"tempat"`
// 	Tanggal_lahir      string    `json:"tanggal_lahir"` // Gunakan string dulu, parsing nanti
// 	Status             string    `json:"status"`
// 	Agama              string    `json:"agama"`
// 	Pekerjaan          string    `json:"pekerjaan"`
// 	Alamat             string    `json:"alamat"`
// 	Kota               string    `json:"kota"`
// 	Kode_pos           int       `json:"kode_pos"`
// 	No_telepon         string    `json:"no_telepon"`
// 	Verified           string    `json:"verified"`
// 	Tanggal_bergabung  string    `json:"tanggal_bergabung"` // Gunakan string dulu, parsing nanti
// 	Email              string    `json:"email"`
// 	Password           string    `json:"password"`
// 	Ktp                string    `json:"ktp"`
// 	Kk                 string    `json:"kk"`
// 	Rekening_koran     string    `json:"rekening_koran"`
// 	Prove_of_residence string    `json:"prove_of_residence"`
// }
