# Category API

API CRUD untuk kategori menggunakan Go dan Gin Framework.

## Fitur

- **GET** `/categories` → Ambil semua kategori
- **POST** `/categories` → Tambah kategori
- **PUT** `/categories/{id}` → Update kategori
- **GET** `/categories/{id}` → Ambil detail satu kategori
- **DELETE** `/categories/{id}` → Hapus kategori

## Menjalankan Lokal

Jalankan dari folder Netlify Function:

```bash
go run netlify/functions/category-api/main.go
```

Atau menggunakan Netlify CLI untuk development:

```bash
netlify dev
```

Server akan berjalan di `http://localhost:8080`

## Deploy ke Netlify

### Opsi 1: Menggunakan Netlify Functions (Go Runtime)

1. Install Netlify CLI:
```bash
npm install -g netlify-cli
```

2. Login ke Netlify:
```bash
netlify login
```

3. Deploy:
```bash
netlify deploy --prod
```

### Opsi 2: Menggunakan Build Command

Netlify akan otomatis mendeteksi Go project dan melakukan build. Pastikan `netlify.toml` sudah dikonfigurasi dengan benar.

## Struktur Project

```
category-api/
├── netlify/
│   └── functions/
│       └── category-api/
│           └── main.go        # Handler untuk Netlify Function & Local Development
├── netlify.toml              # Konfigurasi Netlify
├── go.mod                     # Go dependencies
└── README.md                  # Dokumentasi
```

**Catatan**: File `main.go` di root sudah dihapus. Semua kode ada di `netlify/functions/category-api/main.go` yang bisa dijalankan baik untuk local development maupun sebagai Netlify Function.

## Path yang Tersedia

### Lokal (Development)
- **GET** `http://localhost:8080/categories`
- **POST** `http://localhost:8080/categories`
- **GET** `http://localhost:8080/categories/{id}`
- **PUT** `http://localhost:8080/categories/{id}`
- **DELETE** `http://localhost:8080/categories/{id}`

### Netlify (Production)
Setelah deploy ke Netlify, bisa langsung menggunakan path `/categories` (sama seperti lokal):
- **GET** `https://your-site.netlify.app/categories`
- **POST** `https://your-site.netlify.app/categories`
- **GET** `https://your-site.netlify.app/categories/{id}`
- **PUT** `https://your-site.netlify.app/categories/{id}`
- **DELETE** `https://your-site.netlify.app/categories/{id}`

**Alternatif dengan prefix `/api`** (juga tersedia untuk backward compatibility):
- **GET** `https://your-site.netlify.app/api/categories`
- **POST** `https://your-site.netlify.app/api/categories`
- **GET** `https://your-site.netlify.app/api/categories/{id}`
- **PUT** `https://your-site.netlify.app/api/categories/{id}`
- **DELETE** `https://your-site.netlify.app/api/categories/{id}`

Path akan otomatis di-redirect ke Netlify Function handler.

## Testing

### Lokal - Get All Categories
```bash
curl http://localhost:8080/categories
```

### Lokal - Get Category by ID
```bash
curl http://localhost:8080/categories/1
```

### Lokal - Create Category
```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Category","description":"Test Description"}'
```

### Lokal - Update Category
```bash
curl -X PUT http://localhost:8080/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Category","description":"Updated Description"}'
```

### Lokal - Delete Category
```bash
curl -X DELETE http://localhost:8080/categories/1
```

### Netlify - Get All Categories (langsung /categories)
```bash
curl https://your-site.netlify.app/categories
```

### Netlify - Get Category by ID (langsung /categories)
```bash
curl https://your-site.netlify.app/categories/1
```

### Netlify - Get All Categories (dengan prefix /api)
```bash
curl https://your-site.netlify.app/api/categories
```

### Netlify - Get Category by ID (dengan prefix /api)
```bash
curl https://your-site.netlify.app/api/categories/1
```
