# 🎬 Go Movie API — Production-Grade Folder Structure (Gorilla Mux)

This project is a refactored version of your single-file Gorilla Mux Go application, organized into a **professional production-grade structure** with clear separation of layers (handlers, services, repository, routes, etc.).

---

## 🏗 Folder Structure

```
go-movie-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   └── movie_handler.go
│   ├── models/
│   │   └── movie.go
│   ├── repository/
│   │   └── movie_repo.go
│   ├── routes/
│   │   └── router.go
│   └── services/
│       └── movie_service.go
├── pkg/
│   └── response/
│       └── response.go
├── go.mod
└── README.md
```

---

## 🚀 Quick Start

### 1️⃣ Clone or Download

```bash
git clone https://github.com/shahriar-em0n/CRUD-API-on-GoLang.git
cd CRUD-API-on-GoLang
```

### 2️⃣ Initialize Dependencies

```bash
go mod tidy
```

### 3️⃣ Run the Server

```bash
go run cmd/server/main.go
```

Server will start at:  
👉 **http://localhost:8080**

---

## 🔥 API Endpoints

| Method | Endpoint | Description |
|--------|-----------|-------------|
| GET | `/movie` | Get all movies |
| GET | `/movie/{id}` | Get single movie |
| POST | `/movie` | Create new movie |
| PUT | `/movie/{id}` | Update a movie |
| DELETE | `/movie/{id}` | Delete a movie |

---

## 🧩 Dependencies

- **Gorilla Mux** for routing  
- **Go 1.22+**

Install Gorilla Mux manually (if needed):
```bash
go get github.com/gorilla/mux
```

---

## 🧠 Architecture Overview

**Handler → Service → Repository** Pattern

- **Handler:** Handles HTTP requests/responses  
- **Service:** Contains business logic  
- **Repository:** Handles data storage (in-memory for now)  

---

## 📄 License

MIT License © 2025  
Developed by **Shahriar** 🚀