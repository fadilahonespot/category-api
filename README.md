# Category API

API CRUD untuk mengelola kategori menggunakan Go dan Gin framework.

## Base URL

```
http://localhost:8080
```

## Endpoints

### 1. Health Check

#### GET `/health`

Mengecek status server.

**Response:**
```json
{
  "status": "ok",
  "message": "Server is running",
  "time": "2024-01-15T10:30:00Z"
}
```

**Status Code:** `200 OK`

---

### 2. Get All Categories

#### GET `/categories`

Mengambil semua kategori.

**Response:**
```json
[
  {
    "id": 1,
    "name": "Elektronik",
    "description": "Perangkat elektronik dan peralatan"
  },
  {
    "id": 2,
    "name": "Pakaian",
    "description": "Pakaian dan aksesori pakaian"
  }
]
```

**Status Code:** `200 OK`

**Example:**
```bash
curl http://localhost:8080/categories
```

---

### 3. Get Category by ID

#### GET `/categories/{id}`

Mengambil detail satu kategori berdasarkan ID.

**Parameters:**
- `id` (path parameter, required) - ID kategori

**Response:**
```json
{
  "id": 1,
  "name": "Elektronik",
  "description": "Perangkat elektronik dan peralatan"
}
```

**Status Codes:**
- `200 OK` - Berhasil
- `400 Bad Request` - ID tidak valid
- `404 Not Found` - Kategori tidak ditemukan

**Example:**
```bash
curl http://localhost:8080/categories/1
```

---

### 4. Create Category

#### POST `/categories`

Membuat kategori baru.

**Request Body:**
```json
{
  "name": "Kategori Baru",
  "description": "Deskripsi kategori baru"
}
```

**Request Body Fields:**
- `name` (string, required) - Nama kategori
- `description` (string, optional) - Deskripsi kategori

**Response:**
```json
{
  "id": 6,
  "name": "Kategori Baru",
  "description": "Deskripsi kategori baru"
}
```

**Status Codes:**
- `201 Created` - Kategori berhasil dibuat
- `400 Bad Request` - JSON tidak valid atau name kosong

**Example:**
```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Kategori Baru",
    "description": "Deskripsi kategori baru"
  }'
```

---

### 5. Update Category

#### PUT `/categories/{id}`

Mengupdate kategori yang sudah ada.

**Parameters:**
- `id` (path parameter, required) - ID kategori yang akan diupdate

**Request Body:**
```json
{
  "name": "Kategori Updated",
  "description": "Deskripsi yang diupdate"
}
```

**Request Body Fields:**
- `name` (string, required) - Nama kategori baru
- `description` (string, optional) - Deskripsi kategori baru

**Response:**
```json
{
  "id": 1,
  "name": "Kategori Updated",
  "description": "Deskripsi yang diupdate"
}
```

**Status Codes:**
- `200 OK` - Kategori berhasil diupdate
- `400 Bad Request` - ID tidak valid atau JSON tidak valid
- `404 Not Found` - Kategori tidak ditemukan

**Example:**
```bash
curl -X PUT http://localhost:8080/categories/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Kategori Updated",
    "description": "Deskripsi yang diupdate"
  }'
```

---

### 6. Delete Category

#### DELETE `/categories/{id}`

Menghapus kategori berdasarkan ID.

**Parameters:**
- `id` (path parameter, required) - ID kategori yang akan dihapus

**Response:**
Tidak ada response body.

**Status Codes:**
- `204 No Content` - Kategori berhasil dihapus
- `400 Bad Request` - ID tidak valid
- `404 Not Found` - Kategori tidak ditemukan

**Example:**
```bash
curl -X DELETE http://localhost:8080/categories/1
```

---

## Data Models

### Category

```json
{
  "id": 1,
  "name": "Elektronik",
  "description": "Perangkat elektronik dan peralatan"
}
```

**Fields:**
- `id` (integer) - ID unik kategori (auto-generated)
- `name` (string, required) - Nama kategori
- `description` (string, optional) - Deskripsi kategori

---

## Error Responses

Semua error mengembalikan JSON dengan format:

```json
{
  "error": "Error message"
}
```

**Common Error Status Codes:**
- `400 Bad Request` - Request tidak valid
- `404 Not Found` - Resource tidak ditemukan
- `500 Internal Server Error` - Server error

---

## Menjalankan Lokal

```bash
# Install dependencies
go mod download

# Run server
go run main.go
```

Server akan berjalan di `http://localhost:8080`

---

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

---

## Testing dengan cURL

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

---

## Technology Stack

- **Language**: Go 1.25.0
- **Framework**: Gin Web Framework

---

## License

MIT
