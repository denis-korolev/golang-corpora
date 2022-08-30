help:
	@echo "\
Usage: \n\
    make \n\
         | init                         Поднять приложение с нуля\n\
         | up                           Запустить установленные контейнеры \n\
         | down                         Остановить и уничтожить все контейнеры приложения\n\
         | restart                      Перезапустить проект\n\
         | run-main-go                  Запуск main.go\n\
         | test                         Запуск всех тестов\n\
    "

init: docker-down-clear \
	docker-pull docker-build up
up: docker-up
down: docker-down
restart: down up

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down --remove-orphans

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-pull:
	docker-compose pull --include-deps

docker-build:
	docker-compose build

run-main-go:
	docker-compose run app go run main.go

test:
	docker-compose run app go test -v ./tests/