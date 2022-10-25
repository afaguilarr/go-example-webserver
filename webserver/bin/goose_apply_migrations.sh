cd /app/db_migrations
go run github.com/pressly/goose/v3/cmd/goose postgres "host=postgres user=$1 password=$2 dbname=hello_world sslmode=disable" up
