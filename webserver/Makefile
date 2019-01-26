start:
	gin --appPort 8000 --bin .cache/gin-bin go main.go

docker-up:
	docker-compose up

docker-stop:
	docker-compose stop

db-fixtures:
	go run utils/fixtures.go
