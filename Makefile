run:
	docker-compose up -d --force-recreate --build

test:
	go test ./...

lint:
	gofumpt -l -w .
	golangci-lint run --timeout=5m

generate:
	go generate ./...
	protoc -I=./api/proto --go_out=./internal/delivery/grpc --go_opt=paths=source_relative --go-grpc_out=./internal/delivery/grpc --go-grpc_opt=paths=source_relative ./api/proto/hasher.proto
