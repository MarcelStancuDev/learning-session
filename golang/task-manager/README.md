# Golang Microservices-based Task Manager

## A complete task management application using:

1.  Go (Golang) for backend microservices
2.  RESTful APIs with Gorilla Mux or Gin
3.  PostgreSQL for database
4.  Docker & Docker Compose for containerization
5.  gRPC (optional advanced part)
6.  Redis for caching or session handling
7.  React (optional) for frontend

## 🗂 Project Structure (File Tree)
task-manager/
│
├── api-gateway/
│   ├── main.go
│   ├── go.mod
│   └── handlers/
│       └── task_handler.go
│
├── task-service/
│   ├── main.go
│   ├── go.mod
│   ├── models/
│   │   └── task.go
│   ├── handlers/
│   │   └── task_handler.go
│   ├── db/
│   │   └── postgres.go
│   └── routes/
│       └── router.go
│
├── user-service/
│   ├── main.go
│   ├── go.mod
│   ├── models/
│   │   └── user.go
│   ├── handlers/
│   │   └── user_handler.go
│   ├── db/
│   │   └── postgres.go
│   └── routes/
│       └── router.go
│
├── docker-compose.yml
├── README.md
└── .env