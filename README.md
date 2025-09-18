## Main dependencies

```
go get -u github.com/gin-gonic/gin # Gin Backend Framework
go get -u gorm.io/gorm # GORM for working with PostgreSQL Database
go get -u gorm.io/driver/postgres # GORM driver for PostgreSQL
go get github.com/joho/godotenv # For loading env variables in Golang
```

## Hot reload

The project uses air for hot reload.
https://github.com/air-verse/air

1. Install air
2. Run `air init`
3. Then `air`

## Migrations

- Everytime you create a model or change one: `go run migrate/migrate.go`

## Local db

- Run `docker compose up -d`

## Go version

go1.25.0 linux/amd64