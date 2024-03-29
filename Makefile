help:
	@echo "\
Usage: \n\
    make \n\
         | init                         Поднять приложение с нуля\n\
         | up                           Запустить установленные контейнеры \n\
         | down                         Остановить и уничтожить все контейнеры приложения\n\
         | restart                      Перезапустить проект\n\
         | run-web                  	Запуск web приложения\n\
         | test                         Запуск всех тестов\n\
         | run-cli-list                 Вывод списка cli команд\n\
         | run-shell                 	Запустить контейнер app\n\
         | swagger-generate             Генерировать Swagger\n\
         | update-db             		Обновить данные БД\n\
    "

init: docker-down-clear \
	docker-pull docker-build update-db up
up: run-web
down: docker-down
restart: down up

#docker-up:
#	docker-compose up -d

docker-down:
	docker-compose down --remove-orphans

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-pull:
	docker-compose pull --include-deps

docker-build:
	docker-compose build

run-web:
	docker-compose run --service-ports app go run main.go

run-cli-list:
	docker-compose run --rm app go run bin/main.go

swagger-generate:
	docker-compose run --rm app swag init

run-shell:
	docker-compose run app

test:
	docker-compose run --rm app gotestsum --format testname

update-db:
	docker-compose run --rm app go run bin/main.go importLemma