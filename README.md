# 🎬 Go Movie REST API (with Gorilla Mux)

A simple RESTful API built using **Go (Golang)** and the **Gorilla Mux** router.  
This project demonstrates basic CRUD operations (Create, Read, Update, Delete) for managing movies.

---

## 🚀 Features
- ✅ GET all movies
- ✅ GET a single movie by ID
- ✅ POST (create) a new movie
- ✅ PUT (update) an existing movie
- ✅ DELETE a movie by ID
- ⚙️ Uses `gorilla/mux` for routing
- 🧠 JSON encoding/decoding for clean API responses

---

## 🧩 Tech Stack
- **Language:** Go (Golang)
- **Framework:** Gorilla Mux
- **Libraries:** encoding/json, net/http, log, fmt

---

## 🛠️ Installation & Setup

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/yourusername/go-movie-api.git
cd go-movie-api
```

### 2️⃣ Initialize Go Module
```bash
go mod init go-movie-api
```

### 3️⃣ Install Dependencies
```bash
go get github.com/gorilla/mux
go mod tidy
```

### 4️⃣ Run the Server
```bash
go run main.go
```

Server will start at:
```
http://localhost:8080
```

---

## 🧠 API Endpoints

| Method | Endpoint | Description |
|--------|-----------|-------------|
| GET | `/movies` | Get all movies |
| GET | `/movies/{id}` | Get a movie by ID |
| POST | `/movies` | Create a new movie |
| PUT | `/movies/{id}` | Update an existing movie |
| DELETE | `/movies/{id}` | Delete a movie |

### 🧾 Example JSON Body for POST/PUT:
```json
{
  "isbn": "438227",
  "title": "Inception",
  "director": {
    "firstname": "Christopher",
    "lastname": "Nolan"
  }
}
```

---

## 📦 Example Response
```json
[
  {
    "id": "1",
    "isbn": "438227",
    "title": "Movie One",
    "director": {
      "firstname": "Shahriar",
      "lastname": "Stranger"
    }
  }
]
```

---

## 🧠 Author
**Shahriar** — Go Developer & Problem Solver 🚀

---

## 🪪 License
This project is open source and available under the [MIT License](LICENSE).
