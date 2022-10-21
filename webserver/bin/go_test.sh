go test ./... -cover -coverprofile=c.out
mkdir ./src/report
go tool cover -html=c.out -o src/report/coverage.html
