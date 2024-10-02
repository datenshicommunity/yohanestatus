package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() error {
	// Konfigurasi koneksi database
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return fmt.Errorf("variabel lingkungan database tidak lengkap")
	}

	// Membuat string koneksi
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Membuka koneksi ke database
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi database: %v", err)
	}

	// Memeriksa koneksi
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("gagal melakukan ping ke database: %v", err)
	}

	log.Println("Berhasil terhubung ke database MySQL")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Koneksi database ditutup")
	}
}
