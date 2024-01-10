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

init-config-local:
	if [ ! -f "./internal/adapter/infra/config/configs.yaml" ]; then cp ./internal/adapter/infra/config/configs.yaml.sample ./internal/adapter/infra/config/configs.yaml; fi

start-local: init-config-local
	docker-compose -f docker/docker-compose.yaml up
