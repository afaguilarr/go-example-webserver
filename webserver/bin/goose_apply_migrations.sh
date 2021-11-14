cd /app/db_migrations
goose postgres "host=postgres user=$1 password=$2 dbname=hello_world sslmode=disable" up
