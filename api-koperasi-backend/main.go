package main

import (
	// "api-koperasi-backend/user/shared"
	// controllerProfile "api-koperasi-backend/menu-profile/contProfile"
	"api-koperasi-backend/user/controller"
	// "database/sql"

	"api-koperasi-backend/user/model"

	// corsprofile "api-koperasi-backend/menu-profile/cors"
	modprofile "api-koperasi-backend/menu-profile/modProfile"

	"api-koperasi-backend/user/config"

	"fmt"
	// "log"
	"net/http"
	// "github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	// Inisialisasi koneksi ke database
	config.ConnectDatabase()

	// Membuat router menggunakan mux
	router := mux.NewRouter()

	// Endpoint API
	router.HandleFunc("/api/mahasiswa", model.GetMahasiswa).Methods("GET")
	router.HandleFunc("/api/mhsUpdate", model.UpdateMahasiswa).Methods("PUT")
	router.HandleFunc("/api/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/verify", controller.VerifyEmail).Methods("GET")
	// router.HandleFunc("/login", controller.LoginHandler).Methods("GET")
	router.HandleFunc("/login", controller.LoginUser).Methods("POST")

	// MENU PROFILE
	// -------------------------- ADMIN
	router.HandleFunc("/api/dataAnggota", modprofile.GetDataAnggota).Methods("GET")
	router.HandleFunc("/api/dataDaftarAnggota", modprofile.GetDataDaftarAnggota).Methods("GET")
	router.HandleFunc("/api/profileAnggotaDetail", modprofile.GetDataDaftarAnggotaDetail).Methods("GET")

	// -------------------------- ANGGOTA BIASA
	// router.HandleFunc("/api/register2", controllerProfile.SaveAnggota).Methods("POST") //registrasi ke 2
	// http.HandleFunc("/api/anggota/detail", controllerProfile.GetUserDetails)   // GET
	router.HandleFunc("/api/dataIconPerson", modprofile.GetDataIcon).Methods("GET")

	//get tampilan home berdasarkan email yang login
	router.HandleFunc("/api/profile", controller.GetUserProfile).Methods("GET")
	//handler button submit isi form anggota registrasi ke 2
	router.HandleFunc("/api/register2", controller.RegisterUser2).Methods("PUT") //registrasi ke 2
	//dapatkan data detail user berdasakan yang admin klik
	router.HandleFunc("/api/profileDetail", controller.GetUserProfileDetail).Methods("GET")



	//MAHASISWA COBACOBA YAWWWW
	// router.HandleFunc("/api/students", controller.SaveStudent).Methods("POST")
	// router.HandleFunc("/api/students", controller.SaveStudent).Methods("POST")

	// Tambahkan middleware CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8001"}, // Ganti dengan origin frontend Anda
		// AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE","OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	
	// Bungkus router dengan middleware CORS
	handler := c.Handler(router)

	// Bungkus router dengan middleware CORS
	// http.Handle("/", model.EnableCORS(router))
	// http.Handle("/", controller.EnableCORS(router))
	// http.Handle("/penggunamodel", corsprofile.EnableCORSProfile(router))

	fmt.Println("Server running on http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// Jalankan server
	http.ListenAndServe(":8080", handler)
}
