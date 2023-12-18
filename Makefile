run:
	@go run cmd/go-command.go

install:
	@cd $(GOPATH) && go install github.com/swaggo/swag/cmd/swag@latest; \
	go mod tidy

serve-swagger:
	@mv cmd/go-command.go cmd/main.go; \
	swag init --dir cmd; \
	mv cmd/main.go cmd/go-command.go