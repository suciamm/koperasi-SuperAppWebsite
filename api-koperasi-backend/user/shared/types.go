package shared

// import "time"
// // Struct untuk menerima data JSON
type RegisterRequest struct {
	ID              int
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`

	IsVerified bool
	Jabatan    string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Jabatan string `json:"jabatan"`
}

// {
// 	"email_suci": "x@x.x"
// }


type Registrasi2 struct {
	ID				   int		 `json:"id"`
	Nama_lengkap       string    `json:"nama_lengkap"`
	Jenis_kelamin      string    `json:"jenis_kelamin"`
	Tempat             string    `json:"tempat"`
	Tanggal_lahir      string `json:"tanggal_lahir"`
	Status             string    `json:"status"`
	Agama              string    `json:"agama"`
	Pekerjaan          string    `json:"pekerjaan"`
	Alamat             string    `json:"alamat"`
	Kota               string    `json:"kota"`
	Kode_pos           string       `json:"kode_pos"`
	No_telepon         string    `json:"no_telepon"`
	Verified           string    `json:"verified"`
	Tanggal_bergabung  string `json:"tanggal_bergabung"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Ktp                string    `json:"ktp"`
	Kk                 string    `json:"kk"`
	Rekening_koran     string    `json:"rekening_koran"`
	Prove_of_residence string    `json:"prove_of_residence"`
}  
