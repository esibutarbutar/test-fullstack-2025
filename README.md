# test-fullstack-2025
## Langkah Menjalankan

### 1. Soal Faktorial
go run soal1_faktorial.go

### 2. SistemLogin
#### 1. Clone Repository
#### 2. Jalankan Redis

- *Windows:*  
  Jalankan memurai.exe atau redis-server.exe dari folder hasil ekstrak Redis.
- *Linux/WSL:*  
  sh
  redis-server
  

#### 3. Masukkan Data User ke Redis

Buka terminal baru, masuk ke Redis CLI:
redis-cli
Masukkan data user (contoh):
set login_aberto '{"realname":"Aberto Doni Sianturi","email":"adss@gmail.com","password":"f7c3bc1d808e04732adf679965ccc34ca7ae3441"}'

#### 4. Jalankan Aplikasi Go
go run soal2_sisitemLogin.go
Server akan berjalan di http://localhost:3000.

---

#### Cara Test dengan Postman

1. Buka Postman.
2. Buat request baru:
   - *Method:* POST
   - *URL:* http://localhost:3000/login
   - *Body:* pilih raw dan JSON, isi:
     json
     {
       "username": "aberto",
       "password": "123456789"
     }
     
3. Klik *Send*.

Jika berhasil, response:
json
{
  "message": "Login successful",
  "realname": "Aberto Doni Sianturi",
  "email": "adss@gmail.com"
}

Jika gagal, akan muncul pesan error.

---
