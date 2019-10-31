test/run:
	go test ./app/... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

docker/local/down:
	@echo "============ stopping locally ============"
	docker-compose -f resources/docker/docker-compose.yaml down

docker/local/up:
	@echo "============ starting locally ============"
	docker-compose -f resources/docker/docker-compose.yaml up --build

db:
	docker-compose -f resources/docker/docker-compose.yaml up -d postgres-db pgadmin

migration/create:
	migrate create -ext sql -dir app/migrations $(name)

mockgen:
	mockery -dir ./app/modules/$(module) -name $(interface) -output ./app/modules/$(module)/mocks