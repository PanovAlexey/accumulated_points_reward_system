up:
	cd .docker && docker-compose up --build

down:
	cd .docker && docker-compose down

recreate:
	cd .docker && docker-compose down
	cd .docker && docker-compose up -d --build --force-recreate
	cd .docker && docker-compose up -d

logs:
	cd .docker && docker logs reward_go

build:
	go build -o ./build/gophermart ./cmd/gophermart

exec:
	cd .docker && docker-compose exec reward-go-backend bash
