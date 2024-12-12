package confprofile

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)



var DB *sql.DB

func ConnectDatabase() {
    var err error
    DB, err = sql.Open("mysql", "root:root@tcp(localhost:8889)/koperasi_user_manajemen")
    if err != nil {
        log.Fatal("Gagal menghubungkan ke database:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Gagal ping database:", err)
    }

    log.Println("Koneksi ke database berhasil")
}
