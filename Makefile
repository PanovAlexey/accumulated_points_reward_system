up:
	docker-compose --file .docker/docker-compose.yml up --build

down:
	docker-compose --file .docker/docker-compose.yml down

restart:
	docker-compose --file .docker/docker-compose.yml restart

recreate:
	docker-compose --file .docker/docker-compose.yml down
	docker-compose --file .docker/docker-compose.yml up -d --build --force-recreate
	docker-compose --file .docker/docker-compose.yml up -d

logs:
	cd .docker && docker logs reward_go

build:
	go build -o ./build/gophermart ./cmd/gophermart

exec:
	cd .docker && docker-compose exec reward-go-backend bash

test:
	go test ../...

lint:
	golangci-lint run
