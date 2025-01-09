go mod init ex_proj_go

Запуск проекта
Установите зависимости:
go mod tidy
Запустите сервер:
go run cmd/app/main.go


ex_proj_go
├── cmd/
│   └── main.go
├── configs/
│   └── config.yaml
├── internal/
│   ├── handlers/
│   │   └── user.go
│   ├── models/
│   │   └── user.go
│   ├── services/
│   │   └── worker_pool.go
│   └── db/
│       └── database.go
├── pkg/
│   └── logger/
│       └── logger.go
├── go.mod
└── go.sum
