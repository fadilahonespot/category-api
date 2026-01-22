# Category API

API CRUD untuk mengelola kategori menggunakan Go dan Gin framework.

## Endpoints

- **GET** `/categories` → Ambil semua kategori
- **POST** `/categories` → Tambah kategori
- **PUT** `/categories/{id}` → Update kategori
- **GET** `/categories/{id}` → Ambil detail satu kategori
- **DELETE** `/categories/{id}` → Hapus kategori
- **GET** `/health` → Health check endpoint
- **GET** `/ping` → Ping endpoint

## Deploy ke Render

### Setup Keep-Alive untuk Paket Gratis

Render akan mematikan instance gratis setelah 50 detik tidak ada aktivitas. Untuk menjaga server tetap aktif, gunakan salah satu cara berikut:

#### Opsi 1: Menggunakan UptimeRobot (Recommended)

1. Daftar di [UptimeRobot](https://uptimerobot.com/) (gratis)
2. Tambah monitor baru:
   - **Monitor Type**: HTTP(s)
   - **Friendly Name**: Category API Keep-Alive
   - **URL**: `https://your-app-name.onrender.com/health`
   - **Monitoring Interval**: 5 minutes (atau lebih sering jika tersedia)
3. Simpan monitor

#### Opsi 2: Menggunakan cron-job.org

1. Daftar di [cron-job.org](https://cron-job.org/) (gratis)
2. Buat cron job baru:
   - **Title**: Keep-Alive Category API
   - **URL**: `https://your-app-name.onrender.com/health`
   - **Schedule**: Setiap 1-2 menit (`*/2 * * * *`)
3. Simpan cron job

#### Opsi 3: Menggunakan Render Cron Jobs (jika tersedia)

Jika Render menyediakan cron jobs, buat cron job yang memanggil:
```
curl https://your-app-name.onrender.com/health
```
Set schedule setiap 1-2 menit.

### Cara Deploy

1. Push code ke GitHub
2. Di Render dashboard, pilih "New Web Service"
3. Connect repository GitHub Anda
4. Render akan auto-detect Go
5. Set build command: `go build -o server`
6. Set start command: `./server`
7. Deploy!

### Environment Variables

Tidak ada environment variables yang diperlukan untuk saat ini.

## Docker

### Build Image

```bash
docker build -t category-api .
```

### Run Container

```bash
docker run -d -p 8080:8080 --name category-api category-api
```

### Menggunakan Docker Compose

```bash
# Build dan run
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```

### Build untuk Production

```bash
# Build dengan tag
docker build -t category-api:latest .

# Run dengan environment variable PORT
docker run -d -p 8080:8080 -e PORT=8080 --name category-api category-api:latest
```

## Menjalankan Lokal

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## Testing

```bash
# Health check
curl http://localhost:8080/health

# Get all categories
curl http://localhost:8080/categories

# Get category by ID
curl http://localhost:8080/categories/1

# Create category
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Category","description":"Test Description"}'

# Update category
curl -X PUT http://localhost:8080/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Category","description":"Updated Description"}'

# Delete category
curl -X DELETE http://localhost:8080/categories/1
```
