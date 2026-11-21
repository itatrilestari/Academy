# Academy
Sistem yang mendukung kegiatan pembelajaran, termasuk pengelolaan quiz, pengelolaan sistem pembelajaran, dan evaluasi pembelajaran.

## Tech Stack

- **Language:** Go
- **Framework:** Fiber
- **ORM:** GORM
- **Authentication:** JWT (golang-jwt)
- **Database:** PostgreSQL

## Getting Started

Ikuti langkah-langkah berikut untuk menjalankan project di lingkungan lokal.

### 1. Clone Repository

```bash
git clone https://github.com/itatrilestari/Academy.git
cd Academy
```

### 2. Run Project

Jalankan project:

```bash
go run cmd/main.go
```

Secara default, aplikasi dapat diakses melalui:

```text
http://127.0.0.1:8000/
```

## Authentication

Project ini menggunakan **JSON Web Token (JWT)** sebagai mekanisme autentikasi API.

JWT menyediakan dua jenis token:

- **Access Token:** Digunakan untuk mengakses endpoint API yang membutuhkan autentikasi.
- **Refresh Token:** Digunakan untuk mendapatkan access token baru ketika access token kedaluwarsa.

## Dokumentasi API

Untuk mencoba dan menguji endpoint aplikasi, silakan unduh Postman Collection melalui file berikut:

`data/postman/Academy.postman_collection.json`

### Cara Menggunakan

1. Buka Postman, kemudian import file Collection yang telah diunduh.
2. Atur variabel `base_url` dengan value berikut:

   ```text
   http://localhost:8000/api/
   ```

3. Buka folder **Auth**, kemudian jalankan endpoint **Register** untuk mendapatkan access token.
4. Salin access token yang diperoleh dari response API.
5. Masukkan access token pada header Authorization untuk mengakses endpoint yang memerlukan autentikasi dengan format berikut:

   ```http
   Authorization: Bearer <access_token>
   ```

6. Setelah berhasil melakukan autentikasi, endpoint API sudah dapat digunakan dan diuji melalui Postman.

## License

Repository ini dapat digunakan sebagai referensi.
