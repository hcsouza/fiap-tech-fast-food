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

init-config-local:
	if [ ! -f "./internal/adapter/infra/config/configs.yaml" ]; then cp ./internal/adapter/infra/config/configs.yaml.sample ./internal/adapter/infra/config/configs.yaml; fi


start-local: init-config-local
	docker-compose -f docker/docker-compose.yaml up