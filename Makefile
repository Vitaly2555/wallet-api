up:
	docker-compose up -d

down:
	docker-compose down

migrate-up:
	 migrate -path ./migrations -database $$DATABASE_URL up

migrate-down:
	migrate -path ./migrations -database $$DATABASE_URL down
