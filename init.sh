swag init -g cmd/main.go -o docs
wire ./cmd/.
cd cmd
go build  -ldflags="-w -s" -o main