# Go Terminal Komutları

## Proje Başlatma ve Yönetim

| Komut | Açıklama |
|-------|----------|
| `go mod init <modül-adı>` | Yeni Go modülü oluşturur (package.json gibi) |
| `go mod tidy` | Kullanılmayan bağımlılıkları kaldırır, eksikleri ekler |
| `go get <paket>` | Dış paket indirir (npm install gibi) |
| `go get -u <paket>` | Paketi günceller |

## Çalıştırma ve Derleme

| Komut | Açıklama |
|-------|----------|
| `go run .` | Mevcut dizindeki programı derleyip çalıştırır |
| `go run main.go` | Belirtilen dosyayı çalıştırır |
| `go build` | Çalıştırılabilir dosya oluşturur (.exe) |
| `go build -o isim.exe` | Özel isimle derler |
| `go install` | Derler ve $GOPATH/bin'e koyar |

## Test ve Kalite

| Komut | Açıklama |
|-------|----------|
| `go test` | Testleri çalıştırır (*_test.go dosyaları) |
| `go test -v` | Detaylı test çıktısı |
| `go test -cover` | Test coverage gösterir |
| `go fmt ./...` | Tüm dosyaları formatlar |
| `go vet ./...` | Kod analizi yapar (hata bulur) |

## Bilgi ve Yardım

| Komut | Açıklama |
|-------|----------|
| `go version` | Go versiyonunu gösterir |
| `go env` | Go ortam değişkenlerini gösterir |
| `go doc <paket>` | Paket dökümantasyonunu gösterir |
| `go doc fmt.Println` | Spesifik fonksiyon dökümantasyonu |
| `go help <komut>` | Komut hakkında yardım |

## Bağımlılık Yönetimi

| Komut | Açıklama |
|-------|----------|
| `go list -m all` | Tüm bağımlılıkları listeler |
| `go mod download` | Bağımlılıkları indirir |
| `go mod vendor` | Bağımlılıkları vendor klasörüne kopyalar |
| `go mod graph` | Bağımlılık ağacını gösterir |

## Sık Kullanılan Örnekler

```bash
# Yeni proje başlatma
mkdir my-project
cd my-project
go mod init my-project

# Programı çalıştırma
go run .

# Release build (küçük boyut)
go build -ldflags="-s -w" -o app.exe

# Tüm testleri çalıştır
go test ./...

# Kod formatla ve kontrol et
go fmt ./... && go vet ./...
```

## Ortam Değişkenleri

| Değişken | Açıklama |
|----------|----------|
| `GOPATH` | Go workspace dizini |
| `GOROOT` | Go kurulum dizini |
| `GOOS` | Hedef işletim sistemi (windows, linux, darwin) |
| `GOARCH` | Hedef mimari (amd64, arm64) |

```bash
# Farklı platform için derleme (cross-compile)
set GOOS=linux
set GOARCH=amd64
go build -o app-linux
```
