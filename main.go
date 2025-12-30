package main

import (
	"TorScraper/readwrite"
	"TorScraper/scraper"
	"TorScraper/writer"
	"fmt"
	"os"
)

func main() {
	// 1) Log dosyasını aç
	logFile, err := os.OpenFile("logs/scan_report.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Log dosyası açılamadı:", err)
		return
	}
	defer logFile.Close()

	// 1) URL listesini oku
	urls, err := readwrite.TargetRead("targets.yaml")
	if err != nil {
		fmt.Println("dosya okuma hatası:", err)
		return
	}

	// 2) Tor client oluştur
	client, err := scraper.NewClient()
	if err != nil {
		fmt.Println("Tor client oluşturulamadı:", err)
		return
	}

	// 3) Her URL’i sırayla tara
	fmt.Println("Tarama başlıyor...")
	for _, u := range urls {
		fmt.Println("[INFO] Scanning:", u)

		body, err := scraper.TakeURL(client, u)
		if err != nil {
			fmt.Println("[ERR]  Hata:", err)

			// HATA DURUMUNU LOG DOSYASINA YAZ
			logLine := fmt.Sprintf("[ERR]  Scanning: %s -> %s\n", u, err.Error())
			logFile.WriteString(logLine)

			continue
		}

		fmt.Println("[SUCCESS] Veri alındı:", len(body), "byte")

		// BAŞARILI DURUMU LOG DOSYASINA YAZ
		logLine := fmt.Sprintf("[INFO] Scanning: %s -> SUCCESS (%d bytes)\n", u, len(body))
		logFile.WriteString(logLine)

		// HTML'i kaydet
		savedFile, err := writer.SaveHTML(body)
		if err != nil {
			fmt.Println("[ERR] HTML kaydedilemedi:", err)
		} else {
			fmt.Println("[INFO] HTML kaydedildi:", savedFile)
		}
	}

}
