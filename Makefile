PROTO_DIR=api/proto
PROTO_GEN_DIR=internal/api
DOC_DIR=api/rest

DB_URL=postgresql://postgres:postgres@localhost:5555/postgres?sslmode=disable

generate-proto: create-gen-dir
	protoc --go_out=$(PROTO_GEN_DIR) --go-grpc_out=$(PROTO_GEN_DIR) --go_opt=M$(PROTO_DIR)/sso.proto=./ --go-grpc_opt=M$(PROTO_DIR)/sso.proto=./ --go_opt=paths=import --go-grpc_opt=paths=import $(PROTO_DIR)/sso.proto

create-gen-dir:
	mkdir -p $(PROTO_GEN_DIR)

generate-docs: create-docs-dir
	protoc -I $(PROTO_DIR) --doc_out=$(DOC_DIR) --doc_opt=html,index.html $(PROTO_DIR)/sso.proto

create-docs-dir:
	mkdir -p $(DOC_DIR)

create-db-image:
	docker run --name cifra-sso -p 5555:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

migrate-up:
	migrate -path internal/db/migration -database $(DB_URL) -verbose up

migrate-down:
	migrate -path internal/db/migration -database $(DB_URL) -verbose down

generate-sqlc:
	sqlc generate

build-server:
	go build -o main cmd/sso/main.go

run-server: build-server
	go run cmd/sso/main.go