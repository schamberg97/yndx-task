# Тестовое задание

Serverless-версии находятся в соответствующих папках (goServerless и т.д.). Идентификатор облака - b1geog7ofdtoj6e7fttt ; Доступ предоставлен ya-bkht@yandex.ru (без yandex.ru не работает). Именно serverless-версии и задеплоены (как Yandex Cloud Functions). Все из них выдают UNIX Timestamp.

URL-ы (функции публично доступны):

- https://functions.yandexcloud.net/d4e51lt3fsj5veejuahh
- https://functions.yandexcloud.net/d4e6q0a2bm7knun0pt0j
- https://functions.yandexcloud.net/d4ei2u5osankn87ssgcd

Есть также "обычные" версии в виде Node.js скрипта, Golang-скрипта, bash-скрипта. Так же выполнены скрипты для проверки времени. Во всех случаях, когда аргумент не указан, все скрипты предполагают порт 9000. Упаковывать как докер контейнеры не стал, т.к. было интереснее поиграться с Cloud Functions.



## Node.js

USAGE:

```bash
node ./nodeServer/server.js [PORT]
```

, где аргумент после адреса скрипта - порт, на котором http-сервер должен слушать запросы

Проверка:

```bash
node ./nodeTest/test.js [PORT]
```

## Golang

USAGE:

```bash
cd ./goServer && go run server.go [PORT]
```

, где аргумент [PORT] - порт, на котором http-сервер должен слушать запросы

Проверка:

```bash
cd ./goTest && go run test.go [PORT]
```

## Bash

Очень плохой вариант, потому что:

- постоянный запуск nc и микрообрывы
- плохая и неполная реализация HTTP протокола

USAGE:

```bash
./bashServer/server.sh [PORT]
```

Проверка (по сути просто скрипт, запускающий `curl -v` на указанный порт или порт 9000)

```bash
./bashTest/test.sh [PORT]
```

