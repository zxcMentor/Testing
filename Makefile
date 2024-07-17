include ./env/.env
#  - make build - для сборки приложения;
#  - make test - для запуска unit-тестов;
#  - make docker-build - для сборки Docker-образа с приложением;
#  - make run - для запуска приложения;
#  - make lint - для запуска линтера;

docker-build:
	docker-compose build

run:
	docker-compose up

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml
