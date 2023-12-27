run:
	@go run cmd/go-command.go

install:
	@cd $(GOPATH) && go install github.com/swaggo/swag/cmd/swag@latest; \
	cd $(GOPATH) && go install github.com/vektra/mockery/v2@latest; \
	go mod tidy

tests:
	@go mod tidy; \
	go test ./...

serve-swagger:
	@swag init -g cmd/go-command.go --parseDependency

generate-mocks:
	@mockery --all --keeptree --output tests/mocks