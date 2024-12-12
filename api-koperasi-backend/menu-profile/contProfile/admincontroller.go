package controllerProfile



// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
//     "html/template"

// 	"api-koperasi-backend/menu-profile/model"
// )

// // API untuk mendapatkan semua anggota dalam format JSON
// func GetAllAnggotaAPI(w http.ResponseWriter, r *http.Request) {
//     // Dapatkan data anggota dari model
//     anggotaList, err := model.GetAllAnggota()
//     if err != nil {
//         log.Println("Error fetching anggota data:", err)
//         http.Error(w, "Unable to fetch anggota data", http.StatusInternalServerError)
//         return
//     }

//     // Set header response untuk JSON
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(anggotaList)
// }

// func AdminReadAnggotaHandler(w http.ResponseWriter, r *http.Request) {
//     anggotaList, err := model.GetAllAnggota()
//     if err != nil {
//         log.Println("Error fetching anggota data:", err)
//         http.Error(w, "Unable to load anggota data", http.StatusInternalServerError)
//         return
//     }

//     // Load template
//     tmpl, err := template.ParseFiles("views/sp/admin/daftar_anggota.html")
//     if err != nil {
//         log.Println("Error loading template:", err)
//         http.Error(w, "Unable to load template", http.StatusInternalServerError)
//         return
//     }

//     // Kirim data ke template
//     tmpl.Execute(w, anggotaList)

// }