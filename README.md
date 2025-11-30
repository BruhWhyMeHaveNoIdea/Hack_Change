```markdown
# Hack_Change — Learning Platform (Mono-Repo)

Учебная платформа с микросервисной архитектурой для управления курсами, заданиями и прогрессом учащихся.

## 🚀 Быстрый старт

### Предварительные требования
- Docker & Docker Compose
- Node.js 18+ (для локальной разработки фронтенда)
- Go 1.19+ (для локальной разработки assignment-service)
- Python 3.9+ (для локальной разработки gateway)

### Запуск проекта

1. **Клонируйте и настройте окружение**
```bash
# Копируем файл окружения
cp .env.example .env

# Для Windows PowerShell:
# Copy-Item .env.example .env
```

2. **Запустите все сервисы**
```bash
docker-compose up -d
```

3. **Проверьте работоспособность**
- 🖥️ Frontend: http://localhost:3000
- 🔌 Gateway: http://localhost:8000/health
- 📚 Assignment Service: http://localhost:8080
- 📖 Swagger UI: http://localhost:5000
- 🗄️ Postgres: localhost:5432

## 🏗️ Архитектура системы

```
+----------------+       +----------------+       +-------------------+
|   Frontend     |       |    Gateway     |       |  Assignment Service|
|   (Next.js)    | ----> |   (Python)     | ----> |       (Go)        |
|   :3000        |       |    :8000       |       |      :8080        |
+----------------+       +-------+--------+       +---------+---------+
                                        |                   |
                                        |                   |
                            +-----------v--------+   +------v---------+
                            |  Courses Service   |   |   Postgres DB  |
                            |     (PHP)          |   |                |
                            +--------------------+   +----------------+
```

## 📦 Сервисы

### assignment-service (Go)
**Основной сервис обучения** - API для заданий, прогресса, загрузки решений и авторизации.
- Порт: 8080
- Swagger: http://localhost:5000
- Основные endpoints:
  - `POST /auth/login` - аутентификация
  - `POST /upload` - загрузка решений
  - `GET /progress` - прогресс пользователя
  - `GET /progress/summary` - сводка прогресса

### gateway (Python)
**Публичный API gateway** - роутинг и агрегация запросов к внутренним сервисам.
- Порт: 8000
- Health check: http://localhost:8000/health

### frontend (Next.js)
**Пользовательский интерфейс** - веб-приложение для студентов.
- Порт: 3000 (dev)
- Страницы аутентификации, курсов, заданий

### courses-service (PHP)
**Сервис контента** - управление курсами и учебными материалами.

### База данных (PostgreSQL)
- Порт: 5432
- Миграции: автоматически применяются при запуске

## 🔧 Разработка

### Локальный запуск сервисов

**Frontend (Next.js)**
```bash
cd frontend
npm install
npm run dev
# http://localhost:3000
```

**Assignment Service (Go)**
```bash
cd assignment-service
go run ./cmd/main
# http://localhost:8080
```

**Gateway (Python)**
```bash
cd gateway
python -m venv .venv
# Linux/Mac:
source .venv/bin/activate
# Windows:
# .\.venv\Scripts\Activate
pip install -r requirements.txt
python main.py
```

### Миграции базы данных

Миграции запускаются автоматически. Для ручного запуска:

```bash
# Через Docker Compose
docker-compose run --rm migrate

# Локально (если установлен migrate)
migrate -path ./migrations -database "$DATABASE_URL" up
```

## 🔐 Аутентификация

### Получение тестового JWT токена

1. Используйте seed-данные пользователя:
   - Email: `ivan@example.com`
   - Password: `password123`

2. Получите токен:
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "ivan@example.com", "password": "password123"}'
```

Используйте полученный токен в заголовке:
```
Authorization: Bearer <your_token>
```

## ⚙️ Конфигурация

Основные переменные окружения (`.env`):

```env
# База данных
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=learning_platform
DATABASE_URL=postgres://postgres:password@postgres:5432/learning_platform?sslmode=disable

# JWT
JWT_SECRET_KEY=your-secret-key
JWT_EXPIRE_HOURS=24

# Сервисы
CONTENT_SVC_URL=http://courses-service:80
ASSIGNMENT_SVC_URL=http://assignment-service:8080
GATEWAY_PORT=8000
```

## 🐛 Отладка

### Просмотр логов
```bash
# Логи конкретного сервиса
docker-compose logs -f assignment-service
docker-compose logs -f gateway
docker-compose logs -f frontend

# Все логи
docker-compose logs -f
```

### Проверка базы данных
```bash
# Проверить доступность
docker-compose exec postgres pg_isready

# Подключиться к БД
docker-compose exec postgres psql -U postgres -d learning_platform
```

## 📁 Структура проекта

```
hack_change/
├── assignment-service/     # Go-сервис (основная логика)
│   ├── cmd/
│   ├── internal/
│   └── pkg/
├── gateway/               # Python gateway
├── frontend/              # Next.js приложение
├── courses-service/       # PHP-сервис контента
├── migrations/            # SQL миграции и seed-данные
└── docker-compose.yml
```

## 👥 Команда разработки

- **Александр Ермаков** - Frontend/Backend
- **Дмитрий Карпов** - Frontend  
- **Эдуард Брага** - Backend
- **Антон Лещев** - Backend

## 🚧 Планы развития

- [ ] Добавить CI/CD пайплайн
- [ ] Настроить docker-compose.override.yml для разработки
- [ ] Добавить докеризированный dev-frontend с проксированием
- [ ] Расширить тестовое покрытие
- [ ] Добавить мониторинг и логирование

---

**Hack_Change Learning Platform** - современное решение для онлайн-обучения с микросервисной архитектурой.
```
