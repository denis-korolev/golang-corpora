# golang-corpora
тестовый проект на Golang, чтобы понять как работают основные части Golang.


TODO 
1. `[v]`Сделать так, чтобы каждый раз заново не подтягивались библиотеки https://stackoverflow.com/questions/64400588/go-app-dockerfile-always-downloads-the-modules-on-restart
2. `[v]`Разделить структуры, чтобы не одна большая структура, а несколько мелких
3. XML открывать не сразу весь, а кусками, не все в память грузить. Пакет, который открывает файл и отправляет в канал
4. Передавать в каналы не строки, а структуры
5. Перевести вставку в эластик на балк https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html https://github.com/olivere/elastic/blob/d5049aade9de/recipes/bulk_insert/bulk_insert.go
6. Подключить CLI либу https://github.com/spf13/cobra/blob/main/user_guide.md
7. Подключить ENV 