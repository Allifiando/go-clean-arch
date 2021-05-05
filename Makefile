.PHONY: compile-auth
compile-auth: ## Compile the proto file auth service.
	protoc -I src/pkg/proto/auth/ src/pkg/proto/auth/auth.proto --go_out=plugins=grpc:src/pkg/proto/auth/
 
.PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server server/main.go
	bin/server

.PHONY: run-dev
run-dev:
	bash -c "export ENV=local && nodemon --exec go run src/main.go --signal SIGTERM"