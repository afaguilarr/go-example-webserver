cd /app/db_migrations
go run github.com/pressly/goose/v3/cmd/goose postgres "host=$1 user=$2 password=$3 dbname=$4 sslmode=disable" down
