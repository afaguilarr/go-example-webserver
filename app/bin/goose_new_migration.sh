cd /app/db_migrations
go run github.com/pressly/goose/v3/cmd/goose create $1 sql
