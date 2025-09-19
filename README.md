## Go version

go1.25.0 linux/amd64

## Main dependencies

```
go get -u github.com/gin-gonic/gin # Gin Backend Framework
go get -u gorm.io/gorm # GORM for working with PostgreSQL Database
go get -u gorm.io/driver/postgres # GORM driver for PostgreSQL
go get github.com/joho/godotenv # For loading env variables in Golang
go get golang.org/x/crypto/bcrypt # bcrypt for password hashing
go get github.com/golang-jwt/jwt/v5 # JWT for json web tokens
```

## Hot reload

The project uses air for hot reload.
https://github.com/air-verse/air

1. Install air
2. Run `air init`
3. Then `air`

## Migrations

- Everytime you create a model or change one: `go run migrate/migrate.go`

## Build docker image

```
docker build -t gym-tracker-server .
```

## Run Local db

```
docker compose -f docker-compose.dev.yml up -d
```

## Run prd version

```
docker compose -f docker-compose.prd.yml up -d
```

