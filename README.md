
# Technical Test
Projek ini di buat untuk mengerjakan soal test backend developer dari Bank Ina sebagai berikut

## 1. Membuat Migration Table untuk Users dan Tasks
Saya telah membuat migration tabel mysql di `server/cmd/inits/migration.go` yang sebagai berikut:

*Tabel Users*
- `id` (integer, primary key, auto increment)
- `name` (string, max 255 characters)
- `email` (string, max 255 characters, unique)
- `password` (string, max 255 characters)
- `created_at` (datetime, with current timestamp)
- `updated_at` (datetime, with current timestamp on update)
  
*Tabel Tasks*
- `id` (integer, primary key, auto increment)
- `user_id` (foreign key to users)
- `title` (string, max 255 characters)
- `description` (text)
- `status` (string, max 50 characters, default 'pending')
- `created_at` (datetime, with current timestamp)
- `updated_at` (datetime, with current timestamp on update)

## 2. Membuat API CRUD untuk Users dan Tasks
Saya telah membuat Restful API dengan endpoint berikut
### API Role `/user`
| Route                 | HTTP   | Description                        |
| --------------------- | ------ | ---------------------------------- |
| /                     | POST   | Route used to create user          |
| /                     | GET    | Route used to list user            |
| /:id                  | GET    | Route used to view user            |
| /:id                  | PUT    | Route used to update user          |
| /:id                  | DELETE | Route used to delete user          |

### API Role `/task`
| Route                 | HTTP   | Description                        |
| --------------------- | ------ | ---------------------------------- |
| /                     | POST   | Route used to create task          |
| /                     | GET    | Route used to list task            |
| /:id                  | GET    | Route used to view task            |
| /:id                  | PUT    | Route used to update task          |
| /:id                  | DELETE | Route used to delete task          |

## 3. Menambahkan Autentikasi OAuth 2 untuk Endpoint Task
Saya telah membuat sebuah Autentikasi OAuth 2 berbasis JWT yang dimana untuk mengakses sebuah endpoint `/task` harus memberikan token yang valid supaya bisa akses endpoint task tersebut

## 4. Menggunakan Framework Gin Gonic
Projek test ini di buat dengan teknologi sebagai berikut
### Technologies - Libraries
- ✔️ **[`gin-gonic/gin`](https://github.com/gin-gonic/gin)** -Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster.
- ✔️ **[`go-gorm/gorm`](https://github.com/go-gorm/gorm)** - The fantastic ORM library for Go, aims to be developer friendly
- ✔️ **[`uber-go/zap`](https://github.com/uber-go/zap)** - Blazing fast, structured, leveled logging in Go.
- ✔️ **[`joho/godotenv`](https://github.com/joho/godotenv)**
- ✔️ **[`go-ozzo/ozzo-validation`](https://github.com/go-ozzo/ozzo-validation)** - Ozzo-validation is a Go package that provides configurable and extensible data validation capabilities.
- ✔️ **[`dgrijalva/jwt-go`](https://github.com/dgrijalva/jwt-go)**
- ✔️ **[`crypto/bcrypt`](https://golang.org/x/crypto/bcrypt)**

## 5. Buat Collection Postman
Saya telah membuat meng export collection postmant untuk dokumentasi api atau endpoint apa saja yang telah dibuat atau bisa digunakan, dengan cara mengimport file json tersebut di postmant maka bisa di pakai untuk mencoba projek yang telah saya buat. Untuk file json tersubut ada pada `data/postmant.json`
