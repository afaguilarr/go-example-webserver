go test ./... -cover -coverprofile=c.out
# The -p flag will create the directory if it doesn't exist
mkdir -p ./src/report
go tool cover -html=c.out -o src/report/coverage.html
