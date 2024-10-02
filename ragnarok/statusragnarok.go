package ragnarok

import (
	"fmt"
	"net"
	"os"
	"time"
	"yohanestatus/database"

	"github.com/gofiber/fiber/v2"
)

func HandleRagnarokStatus(c *fiber.Ctx) error {
	// Cek status server menggunakan telnet
	address_ro := os.Getenv("RO_SERVER_ADDRESS")
	if address_ro == "" {
		address_ro = "127.0.0.1" // Menggunakan default jika environment variable tidak diatur
	}
	port_ro := os.Getenv("RO_SERVER_PORT")
	if port_ro == "" {
		port_ro = "6900" // Menggunakan default jika environment variable tidak diatur
	}
	online := checkServerStatus(address_ro, port_ro)
	// Inisialisasi response
	response := struct {
		Online       bool   `json:"online"`
		Players      *struct {
			Online int `json:"online"`
		} `json:"players,omitempty"`
		ServerStatus string `json:"server_status"`
	}{
		Online:       online,
		ServerStatus: "Offline!",
	}

	if online {
		response.ServerStatus = "Online!"

		// Ambil jumlah pemain online dari database
		playersOnline, err := getPlayersOnline()
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Gagal mengambil data pemain: %v", err))
		}

		response.Players = &struct {
			Online int `json:"online"`
		}{
			Online: playersOnline,
		}
	}

	return c.JSON(response)
}

func checkServerStatus(host, port string) bool {
	address := net.JoinHostPort(host, port)
	timeout := time.Second * 5

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

func getPlayersOnline() (int, error) {
	var playersOnline int
	err := database.DB.QueryRow("SELECT COUNT(char_id) AS players_online FROM `char` WHERE `online` > '0'").Scan(&playersOnline)
	if err != nil {
		return 0, err
	}
	return playersOnline, nil
}
