# Film-Library-API

Бэкенд приложение "Фильмотека", которое предоставляет REST API для управления базой данных фильмов.

Для запуска окружения с работающим приложением и СУБД введите команду:

`make run`

Сервер слушает адрес `localhost:8080`

В файле `film-library-api.yaml` предоставлена спецификация на API в формате Swagger 2.0.

Как проверить работу сервера:

- Получить JWT токен администратора по адресу `localhost:8080/authenticate?username=admin&password=admin_password` (Для роли пользователя `username=user`, `password=user_password`). Данные берутся из таблицы `accounts` базы данных.
- Полученный токен необходимо указывать в хедере запросов к серверу, например:
  `curl -X PUT -H "X-Auth-Token: jwtToken" -H "Content-Type: application/json" -d '{"id":1,"year":1920}' http://localhost:8080/films` - здесь `jwtToken` нужно заменить на полученный ранее токен. 
- В GET запросах по маршрутам `/films` и `/actors` используются GET-параметры `limit` и `offset`, по умолчанию они равны **5** и **0** соответственно. Также для фильмов доступны параметры: `search` для поиска фильма по фрагменту названия, по фрагменту имени актёра; `sort` для сортировки по полю (`name`, `rating`, `year`); `order` для выбора способа сортировки (`asc`, `desc`).
- Пример запроса для добавления фильма (в `actors` передаются **id** актёров из таблицы **actors**):
  `curl -X POST -H "X-Auth-Token: jwtToken" -H "Content-Type: application/json" -d '{"name":"Dune: Part Two","year":2024,"description":"sci-fi, action, drama, adventure","rating":8.6,"actors":[2,3]}' http://localhost:8080/films`
 - Пример запроса для добавления актёра:
  `curl -X POST -H "X-Auth-Token: jwtToken" -H "Content-Type: application/json" -d '{"name":"Matthew Paige Damon","gender":"M","birthday":"1960-10-08"}' http://127.0.0.1:8080/actors`

Детали реализации:

- язык реализации - **Go**
- для хранения данных используется **PostgresSQL**
- для логирования используется **стандартный логгер**
- сам http сервер использует стандартную библиотеку **http**
- для авторизации используется **JWT**
- окружение с работающим приложением и СУБД запускается с помощью **docker compose**
- для запуска тестов команда `make test`

## Будущие улучшения

- Вынести настройки БД в файл конфигурации
- Прикрутить более функциональный логгер
- Улучшить покрытие тестами
- Более детальная обработка ошибок

## Схема БД

![schema](images/schema.png)
