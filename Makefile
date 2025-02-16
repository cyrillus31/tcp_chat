.PHONY: client server

client:
	cd client && go run cmd/main.go

server:
	cd server && go run cmd/main.go
