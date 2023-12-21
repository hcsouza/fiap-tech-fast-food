run:
	@go run cmd/go-command.go

install:
	@cd $(GOPATH) && go install github.com/swaggo/swag/cmd/swag@latest; \
	go mod tidy

tests:
	@go mod tidy; \
	go test ./...

serve-swagger:
	@swag init -g cmd/go-command.go --parseDependency