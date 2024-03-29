# Universal Media Organizer (UMO)

[Структура сервиса в Figma](https://www.figma.com/file/nup904QmYhlbg4Aje2v0AM/Untitled?type=design&node-id=0%3A1&mode=design&t=00RI4q4lLIC3gTxb-1)

## Описание
* Сервис предназначен помочь пользователя организовать и отслеживать их любимый медиа-контент. Будь то фильмы, сериалы, книги или даже онлайн-лекции. UMO предоставляет единую платформу для управления всем этим.
* На сервисе пользователи смогут добавлять недостающий контент, отслеживать прогресс по просмотру контента, наблюдать историю и статистику просмотров, оценивать, оставлять комментарии, писать рецензии и самое главное делиться их любимым контентом.
* Также пользователи смогут искать новый контент с помощью страницы с поиском. А благодаря гибкой системе жанров, каждый пользователь сможет найти что-то для себя.

## Frontend
### Запуск
Запустите `npm start` в папке `frontend`. Приложение должно запуститься на `localhost:3001`.

### Технологии
* [React](https://reactjs.org/)
* [React Redux](https://react-redux.js.org/)
* [React Router](https://reactrouter.com/)
* [Redux Persist](https://github.com/rt2zz/redux-persist#readme)
* [Joy-UI](https://mui.com/joy-ui/getting-started/)
* [Axios](https://axios-http.com/)
* [OpenAPI Generator](https://openapi-generator.tech/)

## Backend
* Запустите `./key.sh` в папке `backend` для генерации ключа для JWT.
* Запустите `go run cmd/app/main.go` в папке `backend`. Приложение должно запуститься на `localhost:8000`.

### Технологии
* [Golang](https://golang.org/)
* [Mux](https://github.com/gorilla/mux)
* [JWT](https://jwt.io/)
* [JSend](clevergo.tech/jsend)
* [GORM](https://gorm.io/)
* [Sqlite3](https://www.sqlite.org/index.html)
* [Mockery](https://github.com/vektra/mockery)
* [Testify](https://github.com/stretchr/testify)
