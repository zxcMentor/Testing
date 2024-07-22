#  - make build - для сборки приложения;
#  - make docker-build - для сборки Docker-образа с приложением;
#  - make run - для запуска приложения;
#  - make lint - для запуска линтера;

docker-build:
	docker-compose build

build:
	docker build .

run:
	docker-compose up

lint:
	golangci-lint run ./...
