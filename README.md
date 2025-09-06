# Event Platform

Event Platform — современное веб-приложение для управления событиями, пользователями и подписками на события с использованием Go, GraphQL и MongoDB.

---

## Оглавление

- [Технологии](#технологии)
- [Структура проекта](#структура-проекта)
- [Функциональность](#функциональность)
- [Преимущества](#преимущества)
- [Быстрый старт](#быстрый-старт)
- [Диагностика и отладка](#диагностика-и-отладка)

---

## Технологии

- **Go (Golang):** язык разработки бэкенда.
- **GraphQL (gqlgen):** современный API с queries/mutations/subscriptions.
- **MongoDB:** документно-ориентированная база данных.
- **gqlgen:** генератор кода для GraphQL.
- **Go modules:** управление зависимостями.
- **Docker (опционально):** инфраструктура для быстрой разработки и тестирования.
- **Unit/интеграционное тестирование:** проверки надёжности системы и отдельных компонентов.

---

## Структура проекта

```
├── cmd/ # main (точка входа)
├── graph/ # API-слой GraphQL (модели, схемы, резолверы)
│ ├── model/
│ ├── schema.graphqls
│ ├── resolver.go
│ ├── schema.resolvers.go
│ ├── mutation_resolver.go
│ └── query_resolver.go
├── internal/
│ ├── repository/
│ │ ├── user_repository.go
│ │ └── event_repository.go
│ └── service/
│ ├── user_service.go
│ └── event_service.go
├── go.mod
├── go.sum
├── gqlgen.yml
└── README.md
```
---

## Функциональность

- **Регистрация и управление пользователями:** создание, получение, обновление и удаление пользователей.
- **Создание и просмотр событий:** создание и просмотр информации по событиям.
- **Подписки (subscriptions):** возможность следить за новыми событиями и пользователями в реальном времени.
- **Актуальное хранение данных в MongoDB.**
- **GraphQL Playground:** для тестирования и демонстрации API.

---

## Преимущества

- **Масштабируемость:** Go и MongoDB подходят для нагруженных проектов.
- **Современное API:** GraphQL обеспечивает гибкость и удобство для клиентов любого типа.
- **Разделение слоёв:** Чистая архитектура — есть слои repository (доступ к данным) и service (бизнес-логика).
- **Скорость:** Поддержка real-time подписок.

---

## Быстрый старт

### 1. Клонируйте и установите зависимости:
```
git clone https://github.com/your-username/event-platform.git
cd event-platform
go mod tidy
```

### 2. Запустите MongoDB (локально или через Docker):
```
docker run --name event-mongo -p 27017:27017 -d mongo:6
```

### 3. Запустите приложение:
```
go run ./cmd/event-platform
```

Откройте [http://localhost:8080](http://localhost:8080) для работы с GraphQL Playground.

---

## Диагностика и отладка

- **Ошибки и логи** выводятся в консоль (stdout/stderr).
- **Работа с playground:** можно отлаживать запросы и мутации прямо из браузера.
### Примеры GraphQL-запросов для тестирования

#### 1. Создать пользователя

mutation {
createUser(name: "Alice", email: "alice@example.com", password: "secret") {
id
name
email
createdAt
}
}
text

#### 2. Получить всех пользователей

query {
users {
id
name
email
createdAt
subscriptions {
id
}
}
}
text

#### 3. Создать событие

mutation {
createEvent(
title: "GraphQL Meetup"
description: "Встреча для фанатов GraphQL"
dateTime: "2025-09-07T18:00:00Z"
) {
id
title
description
dateTime
createdAt
}
}
text

#### 4. Получить список событий

query {
events {
id
title
description
dateTime
createdAt
organizer {
name
email
}
}
}
text

#### 5. Получить пользователя по ID

query {
user(id: "ID_ПОЛЬЗОВАТЕЛЯ") {
id
name
email
createdAt
}
}
text

#### 6. Получить событие по ID

query {
event(id: "ID_СОБЫТИЯ") {
id
title
description
dateTime
createdAt
}
}
text

#### 7. Подписаться на событие

mutation {
subscribeToEvent(eventId: "ID_СОБЫТИЯ") {
id
subscriber {
id
name
}
event {
id
title
}
}
}
text

#### 8. Подписаться на пользователя

mutation {
subscribeToUser(userId: "ID_ДРУГОГО_ПОЛЬЗОВАТЕЛЯ") {
id
subscriber {
id
name
}
subscribedToUser {
id
name
}
}
}
text

#### 9. Пример подписки (subscription) на нового пользователя
````
subscription {
subscriber {
id
name
email
createdAt
}
}
```

