docker compose up -d
docker compose exec golang go mod init github.com/mofukuma/golang-gorm-postgres
docker compose exec golang go run migrate.go

docker compose down