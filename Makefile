run:
	@go run internal/main.go

install:
	@cd $(GOPATH) && go install github.com/swaggo/swag/cmd/swag@latest; \
	go mod tidy

serve-swagger:
	@swag init -g cmd/go-command.go --parseDependency