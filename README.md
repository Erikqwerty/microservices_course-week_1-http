# HTTP Проект на Go

Этот проект реализует простой HTTP-сервер и клиент для создания и получения заметок. Сервер использует маршрутизатор `chi` для управления маршрутами и JSON для обмена данными.

---

## Структура проекта

```plaintext
.
├── cmd
│   ├── http_client
│   │   └── main.go           # HTTP-клиент для взаимодействия с сервером
│   └── http_server
│       ├── handler.go        # Обработчики запросов на сервере
│       └── main.go           # Запуск HTTP-сервера
├── models
│   ├── go.mod                # Зависимости проекта
│   ├── go.sum
│   └── README.md             # Документация проекта
```

---

## Установка и настройка

### 1. Клонирование репозитория

```bash
git clone https://github.com/Erikqwerty/microservices_course-week_1-http.git
```

### 2. Установка зависимостей

Проект использует следующие внешние библиотеки:

- `github.com/brianvoe/gofakeit` для генерации случайных данных
- `github.com/go-chi/chi` для маршрутизации
- `github.com/fatih/color` для цветного вывода в консоль

Установите их командой:

```bash
go mod tidy
```

### 3. Запуск сервера

```bash
go run cmd/http_server/main.go
```

Сервер будет запущен по адресу `127.0.0.1:8080`.

### 4. Запуск клиента

В отдельном окне терминала выполните:

```bash
go run cmd/http_client/main.go
```

Клиент создаст новую заметку и попытается её получить.

---

## API Эндпоинты

### 1. Создание заметки

- **URL**: `POST /notes`
- **Описание**: Создаёт новую заметку со случайно сгенерированным содержимым.
- **Тело запроса**:

  ```json
  {
    "title": "string",
    "context": "string",
    "author": "string",
    "is_public": true
  }
  ```

- **Ответ**:
  - `201 Created` при успешном создании
  - JSON-объект с созданной заметкой:

    ```json
    {
      "id": 12345,
      "info": {
        "title": "Название пива",
        "context": "192.168.0.1",
        "author": "Имя автора",
        "is_public": true
      },
      "created_at": "временная метка",
      "updated_at": "временная метка"
    }
    ```

### 2. Получение заметки

- **URL**: `GET /notes/{id}`
- **Описание**: Возвращает заметку по её ID.
- **Ответ**:
  - `200 OK` с данными заметки, если она найдена:

    ```json
    {
      "id": 12345,
      "info": {
        "title": "Название пива",
        "context": "192.168.0.1",
        "author": "Имя автора",
        "is_public": true
      },
      "created_at": "временная метка",
      "updated_at": "временная метка"
    }
    ```

  - `404 Not Found`, если заметка не найдена.

---

## Обзор кода

### Модели

Находятся в пакете `models`:

- **NoteInfo**: Представляет содержимое заметки.
- **Note**: Расширяет `NoteInfo`, добавляя метаданные (ID, временные метки).
- **SyncMap**: Потокобезопасная карта для хранения заметок с использованием мьютекса.

### HTTP-сервер

Находится в `cmd/http_server/main.go` и `cmd/http_server/handler.go`:

- **CreateNote**: Обрабатывает `POST /notes` запросы, создает новый объект `Note` и добавляет его в `SyncMap`.
- **GetNote**: Обрабатывает `GET /notes/{id}`, получает заметку по ID из `SyncMap`.

### HTTP-клиент

Находится в `cmd/http_client/main.go`:

- **CreateNote**: Выполняет `POST` запрос для создания новой заметки со случайными данными.
- **getNote**: Выполняет `GET` запрос для получения заметки по её ID.

---

## Пример использования

### Создание и получение заметок

1. Запустите сервер:

   ```bash
   go run cmd/http_server/main.go
   ```

2. В отдельном терминале запустите клиент для создания и получения заметок:

   ```bash
   go run cmd/http_client/main.go
   ```

   Пример вывода:

   ```plaintext
   Note created:
   {ID:12345 Info:{Title:"Pale Ale" Context:"192.168.1.1" Author:"John Doe" IsPublic:true} CreatedAt:2024-10-07T12:34:56Z UpdatedAt:2024-10-07T12:34:56Z}

   Note info:
   {ID:12345 Info:{Title:"Pale Ale" Context:"192.168.1.1" Author:"John Doe" IsPublic:true} CreatedAt:2024-10-07T12:34:56Z UpdatedAt:2024-10-07T12:34:56Z}
   ```

---

## Основные библиотеки

- **`chi`**: Обеспечивает легковесную маршрутизацию для работы с RESTful маршрутами.
- **`gofakeit`**: Генерирует фейковые данные для тестирования, такие как имена, IP-адреса и т.д.
- **`color`**: Добавляет цвет в вывод терминала для улучшенной читаемости.

---
