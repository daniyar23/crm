.PHONY: run up stop down

up:
	docker-compose up -d

run:
	go run ./main.go

stop:
	docker-compose stop

down:
	docker-compose down
