# GoPigKit
A backend server code repository of [Piggy's Kitchen](https://github.com/wty92911/piggy-s-kitchen) 

## Features
- Lightweight high-performance HTTP routing
- Built-in request and response handling
- Flexible configuration management
- Support for MySQL database and MinIO object storage service

## Directory structure
```txt
.
├── LICENSE
├── README.md
├── build
├── cmd
│   └── pigkit
├── configs
├── internal
│   ├── controllers
│   ├── models
│   ├── pigkit
│   ├── routes
│   └── services
├── pkg
├── tests
└── tmp.out

```
## Installation and configuration
### Prerequisites
- Go 1.22+
- MySQL
- MinIO

### Installation steps
1. Clone the repository
```sh
git clone https://github.com/wty92911/GoPigKit
cd GoPigKit
```
2. Install dependencies
```sh
go mod tidy
```


## Quick start
1. Run database migration
```sh
go run database/migrations/001_create_users_table.sql
```

2. Start the service
```sh
go run cmd/pigkit/main.go
```