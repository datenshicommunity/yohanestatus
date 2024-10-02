package minecraft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type MCServerStatus struct {
	Online  bool `json:"online"`
	Players struct {
		Online int `json:"online"`
	} `json:"players"`
}

type StatusResponse struct {
	Online       bool   `json:"online"`
	Players      *struct {
		Online int `json:"online"`
	} `json:"players,omitempty"`
	ServerStatus string `json:"server_status"`
}

func HandleMCStatus(c *fiber.Ctx) error {
	address_mc := os.Getenv("MC_SERVER_ADDRESS")
	if address_mc == "" {
		address_mc = "127.0.0.1" // Menggunakan default jika environment variable tidak diatur
	}
	
	providers := []struct {
		name string
		url  string
	}{
		{"mcsrvstat", fmt.Sprintf("https://api.mcsrvstat.us/3/%s", address_mc)},
		{"mcstatus", fmt.Sprintf("https://api.mcstatus.io/v2/status/java/%s", address_mc)},
	}

	for _, provider := range providers {
		resp, err := http.Get(provider.url)
		if err != nil {
			continue // Lanjut ke provider berikutnya jika gagal
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue // Lanjut ke provider berikutnya jika gagal membaca respons
		}

		var mcStatus MCServerStatus
		if err := json.Unmarshal(body, &mcStatus); err != nil {
			continue // Lanjut ke provider berikutnya jika gagal mengurai JSON
		}

		response := StatusResponse{
			Online:       mcStatus.Online,
			ServerStatus: "Offline!",
		}

		if mcStatus.Online {
			response.ServerStatus = "Online!"
			response.Players = &struct {
				Online int `json:"online"`
			}{
				Online: mcStatus.Players.Online,
			}
		}

		return c.JSON(response)
	}

	// Jika semua provider gagal
	return c.Status(500).SendString("Gagal mengambil status server dari semua provider")
}
