package controller

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"regexp"
// 	"strings"
	
// 	// "errors"
// 	"koperasi/model"
// 	"net/smtp"

// 	// "Login-regis/config"
// 	// "Login-regis/config"
// 	// "time"
// 	"crypto/rand"
// 	"encoding/hex"

// 	"golang.org/x/crypto/bcrypt"

// 	// "Login-regis/config"
// 	"github.com/gorilla/sessions"
// )

// var store = sessions.NewCookieStore([]byte("secret-key"))

// func RegisterPage(w http.ResponseWriter, r *http.Request) {
// 	tmpl, _ := template.ParseFiles("views/register.html")
// 	tmpl.Execute(w, nil)
// }

// // Struct untuk menyimpan data yang dikirim ke template
// type RegisterPageData struct {
// 	Error string
// }

// // HashPassword meng-hash password menggunakan bcrypt
// func HashPassword(password string) string {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Fatal("Failed to hash password:", err)
// 	}
// 	return string(hashedPassword)
// }

// // USER REGISTRATION HANDLE WITH ERROR HANDLING DIFFERENT PAGE
// // func RegisterUser(w http.ResponseWriter, r *http.Request) {
// //     if r.Method == http.MethodPost {
// //         username := r.FormValue("username")
// //         email := r.FormValue("email")
// //         password := r.FormValue("password")
// //         confirmPassword := r.FormValue("confirm_password")

// //         if password != confirmPassword {
// //             http.Error(w, "Password dan konfirmasi password tidak sama", http.StatusBadRequest)
// //             return
// //         }

// // 		// Hash password
// //         // hashedPassword := HashPassword(password)

// //         //menambahkan kondisi jika user registrasi dengan email yang sama
// //         //start
// //         //cek apakah email yang is_verified nya 0 sudah ada di database
// //         existingUser,  err:= model.GetUserByEmail(email)
// //         if err != nil{
// //             log.Println("Error checking user:", err)
// //             http.Error(w, "Gagal memeriksa data pengguna", http.StatusInternalServerError)
// //             return
// //         }

// //         log.Printf("Existing user data: %+v\n", existingUser)

// //         if existingUser != nil && !existingUser.IsVerified {
// //             // Update data pengguna lama
// //             err := model.UpdateUserByEmail(email, username, HashPassword(password))
// //             if err != nil {
// //                 log.Println("Error updating user:", err)
// //                 http.Error(w, "Gagal memperbarui data pengguna lama", http.StatusInternalServerError)
// //                 return
// //             }
// //         } else if existingUser == nil {
// //             // Buat pengguna baru
// //             err := model.CreateUser(username, email, HashPassword(password))
// //             if err != nil {
// //                 log.Println("Error creating user:", err)
// //                 http.Error(w, "Gagal menyimpan data pengguna", http.StatusInternalServerError)
// //                 return
// //             }
// //         } else {
// //             http.Error(w, "Email sudah digunakan dan terverifikasi", http.StatusBadRequest)
// //             return
// //         }
// //        //end

// //         // // Simpan pengguna dengan status verifikasi false ke database
// //         // err := model.CreateUser(username, email, hashedPassword)
// //         // if err != nil {
// //         //     http.Error(w, "Gagal menyimpan data pengguna", http.StatusInternalServerError)
// //         //     return
// //         // }

// //         // Kirim email verifikasi
// //         verificationToken := generateVerificationToken()
// //         fmt.Println("Generated Token:", verificationToken)

// // 		err = sendVerificationEmail(email, verificationToken)
// //         if err != nil {
// //             http.Error(w, "Gagal mengirim email verifikasi", http.StatusInternalServerError)
// //             return
// //         }

// //         // Simpan token verifikasi di database
// //         err = model.SaveVerificationToken(email, verificationToken)
// //         if err != nil {
// //             http.Error(w, "Gagal menyimpan token verifikasi", http.StatusInternalServerError)
// //             return
// //         }

// //         // Beri tahu pengguna bahwa email verifikasi telah dikirim
// //         fmt.Fprintf(w, "Registrasi berhasil! Silakan cek email Anda untuk verifikasi.")
// //     }
// // }

// func validatePassword(password string) bool {
// 	// Password harus terdiri dari minimal 6 karakter
// 	if len(password) < 6 {
// 		return false
// 	}

// 	// Memeriksa apakah password mengandung huruf besar, huruf kecil, dan karakter khusus
// 	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
// 	hasLower := regexp.MustCompile(`[a-z]`).MatchString
// 	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString // Menggunakan \W untuk karakter non-alphanumeric (termasuk spasi dan simbol)

// 	return hasUpper(password) && hasLower(password) && hasSpecial(password)
// }

// func RegisterUser(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodPost {
// 		// username := r.FormValue("username")
// 		email := r.FormValue("email")
// 		password := r.FormValue("password")
// 		confirmPassword := r.FormValue("confirm_password")

// 		// Validasi password dan konfirmasi password
// 		if password != confirmPassword {
// 			pageVariables := RegisterPageData{
// 				Error: "Password dan konfirmasi password tidak sama.",
// 			}
// 			renderRegisterPage(w, pageVariables)
// 			return
// 		}

// 		//Validasi panjang password
// 		if !validatePassword(password) {
// 			pageVariables := RegisterPageData{
// 				Error: "Password harus terdiri dari minimal 6 karakter, mengandung huruf besar, huruf kecil, dan karakter khusus.",
// 			}
// 			renderRegisterPage(w, pageVariables)
// 			return
// 		}

// 		// Periksa apakah email sudah ada di database
// 		existingUser, err := model.GetUserByEmail(email)
// 		if err != nil {
// 			log.Println("Error checking user:", err)
// 			pageVariables := RegisterPageData{
// 				Error: "Gagal memeriksa data pengguna.",
// 			}
// 			renderRegisterPage(w, pageVariables)
// 			return
// 		}

// 		if existingUser != nil && !existingUser.IsVerified {
// 			// Update data pengguna lama
// 			err := model.UpdateUserByEmail(email, HashPassword(password))
// 			if err != nil {
// 				log.Println("Error updating user:", err)
// 				pageVariables := RegisterPageData{
// 					Error: "Gagal memperbarui data pengguna lama.",
// 				}
// 				renderRegisterPage(w, pageVariables)
// 				return
// 			}
// 		} else if existingUser == nil {
// 			// Buat pengguna baru
// 			err := model.CreateUser(email, HashPassword(password))
// 			if err != nil {
// 				log.Println("Error creating user:", err)
// 				pageVariables := RegisterPageData{
// 					Error: "Gagal menyimpan data pengguna.",
// 				}
// 				renderRegisterPage(w, pageVariables)
// 				return
// 			}
// 		} else {
// 			pageVariables := RegisterPageData{
// 				Error: "Email sudah digunakan dan terverifikasi.",
// 			}
// 			renderRegisterPage(w, pageVariables)
// 			return
// 		}

// 		// Kirim email verifikasi
// 		verificationToken := generateVerificationToken()
// 		err = sendVerificationEmail(email, verificationToken)
// 		if err != nil {
// 			pageVariables := RegisterPageData{
// 				Error: "Gagal mengirim email verifikasi.",
// 			}
// 			renderRegisterPage(w, pageVariables)
// 			return
// 		}

// 		// Simpan token verifikasi di database
// 		err = model.SaveVerificationToken(email, verificationToken)
// 		if err != nil {
// 			pageVariables := RegisterPageData{
// 				Error: "Gagal menyimpan token verifikasi.",
// 			}
// 			renderRegisterPage(w, pageVariables)
// 			return
// 		}

// 		// Registrasi berhasil
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 	} else {
// 		// Kirim data kosong jika pertama kali membuka halaman
// 		renderRegisterPage(w, RegisterPageData{Error: ""})
// 	}
// }

// func GetStaticFile(w http.ResponseWriter, r *http.Request) {
// 	const prefix = "/statics/"
// 	if !strings.HasPrefix(r.URL.Path, prefix) {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	assetName := strings.TrimPrefix(r.URL.Path, prefix)
// 	http.ServeFile(w, r, fmt.Sprintf("statics/%s", assetName))
// }

// func renderRegisterPage(w http.ResponseWriter, pageVariables RegisterPageData) {
// 	tmpl, err := template.ParseFiles("views/register.html")
// 	if err != nil {
// 		http.Error(w, "Unable to load template", http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, pageVariables)
// }

// // GENERATE TOKEN EMAIL
// func generateVerificationToken() string {
// 	token := make([]byte, 32)
// 	_, err := rand.Read(token)
// 	if err != nil {
// 		fmt.Println("berhasil generate")
// 		log.Fatal("Error generating token:", err)
// 	}
// 	return hex.EncodeToString(token)
// }

// // SENDING EMAIL VERIFICATION
// func sendVerificationEmail(to, token string) error {
// 	from := "systexsuci22@gmail.com"
// 	password := "lqocianpnkykmwqq"
// 	smtpServer := "smtp.gmail.com"
// 	port := "587"
// 	// port := "465"
// 	subject := "Email Verifikasi"
// 	body := fmt.Sprintf("Klik link berikut untuk memverifikasi email Anda: http://localhost:8010/verify?token=%s", token)
// 	log.Println("Received token:", token)

// 	msg := []byte("To: " + to + "\r\n" +
// 		"Subject: " + subject + "\r\n" +
// 		"\r\n" + body)

// 	auth := smtp.PlainAuth("", from, password, smtpServer)

// 	// Log untuk membantu debugging
// 	log.Println("Sending email to:", to)
// 	log.Println("SMTP server:", smtpServer+":"+port)

// 	err := smtp.SendMail(smtpServer+":"+port, auth, from, []string{to}, msg)
// 	if err != nil {
// 		log.Println("Error sending email:", err)
// 	}

// 	return smtp.SendMail(smtpServer+":"+port, auth, from, []string{to}, msg)
// }

// // func VerifyEmail(w http.ResponseWriter, r *http.Request) {
// //     token := r.URL.Query().Get("token")
// //     if token == "" {
// //         http.Error(w, "Token tidak ditemukan", http.StatusBadRequest)
// //         return
// //     }

// //     err := model.VerifyUserEmail(token)
// //     if err != nil {
// //         http.Error(w, "Verifikasi gagal", http.StatusInternalServerError)
// //         return
// //     }

// //     fmt.Fprintf(w, "Email Anda berhasil diverifikasi. Silakan login.")
// // }

// // TIMEOUT VERIFIKASI LINKING EMAIL REGISTERED
// func VerifyEmail(w http.ResponseWriter, r *http.Request) {
// 	token := r.URL.Query().Get("token")
// 	if token == "" {
// 		http.Error(w, "Token tidak ditemukan", http.StatusBadRequest)
// 		return
// 	}

// 	err := model.VerifyUserEmail(token)
// 	if err != nil {
// 		if err.Error() == "token sudah kedaluwarsa" {
// 			http.Error(w, "Token sudah kedaluwarsa. Silakan mendaftar ulang.", http.StatusUnauthorized)
// 			return
// 		}
// 		log.Println("Token received for verification:", token)
// 		http.Error(w, "Verifikasi gagal", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "Email Anda berhasil diverifikasi. Silakan login.")
// }

// // RESET PASSWORD AND USERNAME VIA EMAIL SESI 1

// // func ResetPassword(w http.ResponseWriter, r *http.Request) {
// //     if r.Method == http.MethodPost {
// //         email := r.FormValue("email")

// //         // Cek apakah email terdaftar
// //         user, err := model.GetUserByEmail(email)
// //         if err != nil || user == nil {
// //             http.Error(w, "Email tidak ditemukan", http.StatusBadRequest)
// //             return
// //         }

// //         // Generate token untuk reset password
// //         resetToken := generateVerificationToken()

// //         // Simpan token di database
// //         err = model.SaveVerificationToken(email, resetToken)
// //         if err != nil {
// //             http.Error(w, "Gagal menyimpan token reset password", http.StatusInternalServerError)
// //             return
// //         }

// //         // Kirim email dengan tautan reset password
// //         resetLink := fmt.Sprintf("http://localhost:8010/form-reset?token=%s", resetToken)
// //         emailBody := fmt.Sprintf("Klik tautan berikut untuk mengatur ulang password Anda: %s", resetLink)
// //         err = sendVerificationEmail(email, emailBody)
// //         if err != nil {
// //             http.Error(w, "Gagal mengirim email reset password", http.StatusInternalServerError)
// //             return
// //         }

// //         fmt.Fprintf(w, "Email untuk reset password telah dikirim.")
// //     } else {
// //         http.ServeFile(w, r, "views/reset-password.html")
// //     }
// // }

// // func UpdatePassword(w http.ResponseWriter, r *http.Request) {
// //     if r.Method == http.MethodPost {
// //         token := r.FormValue("token")
// //         newPassword := r.FormValue("password")
// //         confirmPassword := r.FormValue("confirm_password")

// //         if newPassword != confirmPassword {
// //             http.Error(w, "Password dan konfirmasi password tidak cocok", http.StatusBadRequest)
// //             return
// //         }

// //         // Verifikasi token
// //         email, err := model.GetEmailByToken(token)
// //         if err != nil {
// //             http.Error(w, "Token tidak valid atau sudah kedaluwarsa", http.StatusBadRequest)
// //             return
// //         }

// //         // Hash password baru
// //         hashedPassword := HashPassword(newPassword)

// //         // Perbarui password pengguna
// //         err = model.UpdatePasswordByEmail(email, hashedPassword)
// //         if err != nil {
// //             http.Error(w, "Gagal memperbarui password", http.StatusInternalServerError)
// //             return
// //         }

// //         // Hapus token setelah digunakan
// //         err = model.DeleteToken(token)
// //         if err != nil {
// //             http.Error(w, "Gagal menghapus token", http.StatusInternalServerError)
// //             return
// //         }

// //         fmt.Fprintf(w, "Password berhasil diperbarui. Silakan login.")
// //     } else {
// //         http.ServeFile(w, r, "views/form-reset.html")
// //     }
// // }

// //---------------------------
// //RESET PASSWORD SESI 2
// //--------------------------------------

// func RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		email := r.FormValue("email")

// 		// Periksa apakah email ada di database
// 		user, err := model.GetUserByEmail(email)
// 		if err != nil {
// 			http.Error(w, "Gagal memeriksa data pengguna", http.StatusInternalServerError)
// 			return
// 		}

// 		if user == nil {
// 			http.Error(w, "Email tidak terdaftar", http.StatusBadRequest)
// 			return
// 		}

// 		// Buat token reset password
// 		resetToken := generateVerificationToken()

// 		// Simpan token reset password
// 		err = model.SaveVerificationToken(email, resetToken)
// 		if err != nil {
// 			http.Error(w, "Gagal menyimpan token reset password", http.StatusInternalServerError)
// 			return
// 		}

// 		// Kirim email reset password
// 		resetLink := fmt.Sprintf("http://localhost:8010/form-reset?token=%s", resetToken)
// 		err = sendResetPasswordEmail(email, resetLink)
// 		if err != nil {
// 			http.Error(w, "Gagal mengirim email reset password", http.StatusInternalServerError)
// 			return
// 		}

// 		fmt.Fprintf(w, "Link reset password telah dikirim ke email Anda.")
// 	}
// }

// func sendResetPasswordEmail(to, resetLink string) error {
// 	from := "systexsuci22@gmail.com"
// 	password := "lqocianpnkykmwqq"
// 	smtpServer := "smtp.gmail.com"
// 	port := "587"
// 	subject := "Reset Password"
// 	body := fmt.Sprintf("Klik link berikut untuk mereset password Anda: %s", resetLink)

// 	msg := []byte("To: " + to + "\r\n" +
// 		"Subject: " + subject + "\r\n" +
// 		"\r\n" + body)

// 	auth := smtp.PlainAuth("", from, password, smtpServer)
// 	return smtp.SendMail(smtpServer+":"+port, auth, from, []string{to}, msg)
// }

// func ShowResetPasswordForm(w http.ResponseWriter, r *http.Request) {
// 	token := r.URL.Query().Get("token")
// 	if token == "" {
// 		http.Error(w, "Token tidak valid", http.StatusBadRequest)
// 		return
// 	}

// 	tmpl, err := template.ParseFiles("views/form-reset.html")
// 	if err != nil {
// 		http.Error(w, "Gagal memuat template", http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl.Execute(w, map[string]interface{}{
// 		"Token": token,
// 	})
// }

// func UpdatePassword(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		token := r.FormValue("token")
// 		password := r.FormValue("password")
// 		confirmPassword := r.FormValue("confirm_password")

// 		if password != confirmPassword {
// 			http.Error(w, "Password dan konfirmasi password tidak cocok", http.StatusBadRequest)
// 			return
// 		}

// 		// Periksa token dan dapatkan email
// 		var email string
// 		err := model.VerifyResetToken(token, &email)
// 		if err != nil {
// 			http.Error(w, "Token tidak valid atau kadaluarsa", http.StatusBadRequest)
// 			return
// 		}

// 		// Hash password baru
// 		hashedPassword := HashPassword(password)

// 		// Perbarui password pengguna
// 		err = model.UpdatePassword(email, hashedPassword)
// 		if err != nil {
// 			http.Error(w, "Gagal memperbarui password", http.StatusInternalServerError)
// 			return
// 		}

// 		fmt.Fprintf(w, "Password berhasil diperbarui!")
// 	}
// }

// // LOGIN
// type LoginPageData struct {
// 	Error string
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	// Render the login page. Initially, no error is passed.
// 	pageVariables := LoginPageData{}
// 	// gbr := "base64xxx"
// 	tmpl, err := template.ParseFiles("views/login.html")
// 	if err != nil {
// 		http.Error(w, "Unable to load template", http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, pageVariables)
// }

// func LoginUser(w http.ResponseWriter, r *http.Request) {
// 	// Handle request
// 	if r.Method == http.MethodPost {
// 		// Parsing form
// 		email := r.FormValue("email")
// 		password := r.FormValue("password")

// 		user, err := model.ValidateUser(email, password)
// 		if err != nil {
// 			log.Println("Login error:", err)

// 			// Render the login page with an error message
// 			pageVariables := LoginPageData{
// 				Error: "email or Password is incorrect", // Set error message if validation fails
// 			}
// 			tmpl, tmplErr := template.ParseFiles("views/login.html")
// 			if tmplErr != nil {
// 				http.Error(w, "Unable to load template", http.StatusInternalServerError)
// 				return
// 			}
// 			tmpl.Execute(w, pageVariables)
// 			return
// 		}

// 		// fmt.Println("sebelum masuk kondisi jabatab")
// 		// Periksa jabatan user dari data yang dikembalikan
// 		switch user.Jabatan {
// 		case "anggota":
// 			// fmt.Println("masuk enggota")
// 			// Redirect ke halaman anggota
// 			http.Redirect(w, r, "/sp/anggota/home", http.StatusSeeOther)
// 		case "admin":
// 			// fmt.Println("masuk admin")
// 			// Redirect ke halaman admin
// 			http.Redirect(w, r, "/sp/admin/home", http.StatusSeeOther)
// 		default:
// 			// Jika jabatan tidak diketahui, beri pesan error
// 			http.Error(w, "Invalid role", http.StatusUnauthorized)
// 		}
// 		//end

// 		// User authenticated, create session
// 		session, _ := store.Get(r, "user-session")
// 		session.Values["userID"] = user.ID
// 		session.Values["email"] = user.Email
// 		session.Values["jabatan"] = user.Jabatan // Tambahkan jabatan ke sesi

// 		session.Save(r, w)

// 		// Redirect to user management page
// 		// http.Redirect(w, r, "/user-mg", http.StatusSeeOther)
// 		// http.Redirect(w, r, "/sp/admin/home", http.StatusSeeOther)

// 	} else {
// 		// If method is not POST, return a method not allowed error
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}

// }

// func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "user-session")
// 	email, ok := session.Values["email"].(string)
// 	if !ok {
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 		return
// 	}

// 	// tmpl, err := template.ParseFiles("views/user-mg.html")
// 	tmpl, err := template.ParseFiles("views/sp/anggota/home.html")

// 	if err != nil {
// 		log.Println("Error parsing template:", err)
// 		http.Error(w, "Internal Server Error- welcome handler controller", http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl.Execute(w, map[string]string{
// 		"Email": email,
// 	})
// }





// // SIMPAN PINJAM ANGGOTA 

// // Fungsi untuk mengambil data pengguna berdasarkan ID
// func ProfileReadHandler(w http.ResponseWriter, r *http.Request) {
//     session, _ := store.Get(r, "user-session")
//     email, ok := session.Values["email"].(string)
//     if !ok {
//         log.Println("Email not found in session")
//         http.Redirect(w, r, "/login", http.StatusSeeOther)
//         return
//     }

//     // Panggil data anggota berdasarkan email
//     anggota, err := model.GetSPAnggotaByEmail(email)
//     if err != nil {
//         log.Println("Error retrieving anggota:", err)
//         http.Error(w, "Data anggota tidak ditemukan", http.StatusInternalServerError)
//         return
//     }

//     tmpl, err := template.ParseFiles("views/sp/anggota/profile.html")
//     if err != nil {
//         log.Println("Error parsing template:", err)
//         http.Error(w, "Internal Server Error profile controller", http.StatusInternalServerError)
//         return
//     }
// 	// fmt.Println("masuk function profilll")
//     tmpl.Execute(w, anggota)
// }

// //END SIMPAN PINJAM ANGGOTA





// //SIMPAN ANGGOTA ADMIN

// func AdminHomeHandler(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "user-session")
// 	email, ok := session.Values["email"].(string)
// 	if !ok {
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 		return
// 	}

// 	// tmpl, err := template.ParseFiles("views/user-mg.html")
// 	tmpl, err := template.ParseFiles("views/sp/admin/home.html")

// 	if err != nil {
// 		log.Println("Error parsing template:", err)
// 		http.Error(w, "Internal Server Error- welcome handler controller", http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl.Execute(w, map[string]string{
// 		"Email": email,
// 	})
// }

// func AdminProfileHandler(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "user-session")
//     email, ok := session.Values["email"].(string)

//     if !ok {
//         http.Redirect(w, r, "/login", http.StatusSeeOther)
//         return
//     }

//     admin, err := model.GetSPAdminByEmail(email)
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









// // Fungsi untuk mengambil data pengguna berdasarkan ID
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





// func LoginUser(w http.ResponseWriter, r *http.Request) {
// 	// Handle request
// 	if r.Method == http.MethodPost {
// 		// Parsing form
// 		email := r.FormValue("email")
// 		password := r.FormValue("password")

// 		user, err := model.ValidateUser(email, password)
// 		if err != nil {
// 			log.Println("Login error:", err)

// 			// Render the login page with an error message
// 			pageVariables := LoginPageData{
// 				Error: "email or Password is incorrect", // Set error message if validation fails
// 			}
// 			tmpl, tmplErr := template.ParseFiles("views/login.html")
// 			if tmplErr != nil {
// 				http.Error(w, "Unable to load template", http.StatusInternalServerError)
// 				return
// 			}
// 			tmpl.Execute(w, pageVariables)
// 			return
// 		}

// 		// fmt.Println("sebelum masuk kondisi jabatab")
// 		// Periksa jabatan user dari data yang dikembalikan
// 		switch user.Jabatan {
// 		case "anggota":
// 			// fmt.Println("masuk enggota")
// 			// Redirect ke halaman anggota
// 			http.Redirect(w, r, "/sp/anggota/home", http.StatusSeeOther)
// 		case "admin":
// 			// fmt.Println("masuk admin")
// 			// Redirect ke halaman admin
// 			http.Redirect(w, r, "/sp/admin/home", http.StatusSeeOther)
// 		default:
// 			// Jika jabatan tidak diketahui, beri pesan error
// 			http.Error(w, "Invalid role", http.StatusUnauthorized)
// 		}
// 		//end

// 		// User authenticated, create session
// 		session, _ := store.Get(r, "user-session")
// 		session.Values["userID"] = user.ID
// 		session.Values["email"] = user.Email
// 		session.Values["jabatan"] = user.Jabatan // Tambahkan jabatan ke sesi

// 		session.Save(r, w)

// 		// Redirect to user management page
// 		// http.Redirect(w, r, "/user-mg", http.StatusSeeOther)
// 		// http.Redirect(w, r, "/sp/admin/home", http.StatusSeeOther)

// 	} else {
// 		// If method is not POST, return a method not allowed error
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}

// }