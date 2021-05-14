all: clean build-server

clean:
	rm -f bin/auth-service

build-server: bin/auth-service

bin/auth-service:
	go build -o bin/auth-service cmd/server/*.go

PHONY: run-server
run-server:
	@go run cmd/server/main.go

PHONY: test-all
test-all:
	@echo "pkg/repository/inmemory" && go test pkg/repository/inmemory/*

fresh:
	@cd cmd/server && fresh
