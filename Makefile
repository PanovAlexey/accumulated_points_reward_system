up:
	cd .docker && docker-compose up --build

down:
	cd .docker && docker-compose down

recreate:
	cd .docker && docker-compose down
	cd .docker && docker-compose up -d --build --force-recreate
	cd .docker && docker-compose up -d
