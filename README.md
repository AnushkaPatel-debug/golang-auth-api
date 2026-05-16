# Golang JWT Authentication API

## Features
- User Signup
- User Login
- JWT Authentication
- Protected Routes
- Role Based Access Control
- Admin Middleware
- SQLite Database
- Password Hashing using bcrypt

---

## Tech Stack
- Golang
- Gin Framework
- GORM
- SQLite
- JWT
- bcrypt

---

## API Endpoints

### Signup
POST /signup

### Login
POST /login

### Profile
GET /profile

Protected Route requiring JWT token.

### Users
GET /users

Admin-only protected route.

---

## Authentication
JWT tokens are used for authentication.

Authorization Header Format:

Bearer your_token_here

---

## Run Project

```bash
go run main.go

server runs on:
http://localhost:8090


golang-auth-api/
│
├── database/
├── handlers/
├── middleware/
├── models/
├── utils/
├── main.go
└── README.md