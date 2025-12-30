# Tor Scraper

Basit bir Go uygulaması olarak tasarlanan bu proje, `.onion` adreslerini Tor ağı üzerinden ziyaret eder, HTML çıktısını kaydeder ve her isteğin sonucunu log dosyasına yazar. Amaç, Tor üzerinden otomatik veri toplama sürecini minimal ve anlaşılır bir yapıda göstermek.

---

## Özellikler

- `targets.yaml` dosyasından hedef adresleri okur  
- Tüm HTTP istekleri Tor Browser’ın SOCKS5 proxy’si üzerinden yapılır (`127.0.0.1:9150`)  
- Başarılı isteklerden dönen HTML verisi `screenshots/` klasörüne timestamp ile kaydedilir  
- Her URL’in sonucu `logs/scan_report.log` dosyasına eklenir  

---

## Gereksinimler

- Go (Golang)
- Tor Browser (arka planda açık olmalı)
- SOCKS5 proxy portu: **9150**

---

## Çalıştırma

1. Tor Browser’ı aç ve arka planda çalışır durumda bırak.
2. Terminalden proje dizinine geç.
3. Uygulamayı başlat:

```bash
go run main.go





