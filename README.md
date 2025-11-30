
Hack_Change — Learning Platform (mono-repo)
Кратко: в этом репозитории — простая учебная платформа с несколькими микросервисами:
·	assignment-service — сервис на Go (API для заданий, прогресса, загрузки решений, авторизации);
·	gateway — Python-прокси/агрегатор (роутинг и публичный API);
·	frontend — Next.js приложение (интерфейс пользователей);
·	courses-service — простой PHP-сервис контента (курсы/страницы);
·	postgres — база данных (контейнер в docker-compose).
В корне проекта есть docker-compose.yml и пример переменных окружения в .env.example.
Что реализовано:
·	Основной набор API для управления курсами, заданиями, прогрессом и загрузками (см. assignment-service).
·	Gateway выполняет проксирование запросов к внутренним сервисам и предоставляет простую точку входа на :8000.
·	Frontend (Next.js) реализует страницы для аутентификации и просмотра курсов/заданий.
·	Миграции и seed-данные для инициализации БД лежат в папке migrations.
Архитектура (упрощённо)
                                                 +----------------+
                                                 |   frontend     |  (Next.js, :3000 при dev)
                                                 +--------+-------+
                                                                    |
                                                                    | HTTPS/REST
                                                                    v
                                                    +-------+-------+
                                                    |    gateway    |  (Python, :8000)
                                                    +---+-------+---+
                            internal calls |       | external proxy
                                                 +---v---+ +--v-------+
                                                 |assign- | | courses |
                                                 |ment-   | | service |
                                                 |service | | (PHP)   |
                                                 | (Go)   | |         |
                                                 +---+----+ +----+----+
                                                         |         |
                                                         |         |
                                                    +--v---------v--+
                                                    |   Postgres DB  |
                                                    +----------------+

Ключевая идея: gateway — публичный вход; assignment-service содержит логику обучения и хранит данные; courses-service даёт контент курса.
Быстрый старт (рекомендуемый)
1.	Скопируйте пример .env:
cp .env.example .env

(На Windows PowerShell используйте Copy-Item .env.example .env).
2.	Отредактируйте .env при необходимости (пароли, ключ JWT и т.п.).
3.	Запустите всю систему через Docker Compose (из корня репозитория):
docker-compose up -d

4.	Проверки после запуска:
·	Postgres: localhost:<POSTGRES_PORT> (по умолчанию 5432).
·	Gateway: http://localhost:8000/health — должен вернуть статус сервиса.
-assignment-service API: http://localhost:8080.
·	Swagger UI: http://localhost:5000 (API documentation для assignment-service).
·	Frontend: http://localhost:3000
Запуск миграций
Миграции автоматически запускаются контейнером migrate в docker-compose.yml. Файлы миграций расположены в папке migrations.
Если хотите запустить миграции вручную:
docker-compose run --rm migrate

Или (если вы используете DATABASE_URL локально):
docker run --rm -v %cd%/migrations:/migrations --network host migrate/migrate:latest -path /migrations -database "%DATABASE_URL%" up

Локальная разработка
·	Frontend (Next.js):
o	Установите зависимости: npm install или pnpm install в папке frontend.
o	Запустите dev сервер: npm run dev (обычно на http://localhost:3000).
·	Assignment service (Go):
o	Перейдите в assignment-service, установите зависимости и запустите:
cd assignment-service
go run ./cmd/main

o	Сервис слушает :8080 (при использовании .env он подключится к БД).
·	Gateway (Python):
o	В папке gateway есть requirements.txt; создайте виртуальное окружение и установите зависимости:
python -m venv .venv; .\.venv\Scripts\Activate; pip install -r requirements.txt
python main.py

Переменные окружения
Основные переменные находятся в .env.example. Коротко:
·	POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB — настройки БД.
·	DATABASE_URL — полная строка подключения (используется мигратором).
·	JWT_SECRET_KEY, JWT_EXPIRE_HOURS — для генерации JWT.
·	CONTENT_SVC_URL, ASSIGNMENT_SVC_URL, GATEWAY_PORT — URL/порты для docker-сетей.
Пример: скопируйте ./.env.example → ./.env и при необходимости отредактируйте.
Как получить тестовый JWT (пример)
1.	В migrations есть seed, который создаёт тестового пользователя ivan@example.com / password123.
2.	Выполните POST запрос:
POST http://localhost:8080/auth/login
Content-Type: application/json

{
    "email": "ivan@example.com",
    "password": "password123"
}

В ответ вы получите JWT — используйте его в Swagger (Authorize) или в заголовке Authorization: Bearer <token>.
Полезные конечные точки (assignment-service)
·	POST /auth/login — логин
·	POST /upload — загрузка решения
·	GET /progress — прогресс по текущему пользователю
·	GET /progress/summary — сводка прогресса по курсам
·	GET /:assignmentId/feedback — фидбэк по отправке
Полный список и схемы доступны в Swagger UI (http://localhost:5000).
Отладка и советы
·	Посмотрите логи контейнеров:
docker-compose logs -f assignment-service
docker-compose logs -f gateway

·	Если миграции не применяются, проверьте переменную DATABASE_URL в .env и доступность Postgres (pg_isready).
·	Для разработки удобно запускать сервисы локально (Go и Node) и подключать их к Docker-сети или локальному Postgres через LOCAL_DATABASE_URL.
Структура репозитория (важные каталоги)
·	assignment-service/ — Go-сервис (handlers, internal, pkg, cmd)
·	gateway/ — Python прокси и маршрутизация
·	frontend/ — Next.js приложение
·	courses-service/ — PHP-контент
·	migrations/ — SQL миграции и seed-скрипты
Дальше/улучшения
·	Добавить CI для запуска миграций и тестов.
·	Организовать docker-compose.override.yml для локальной разработки.
·	Добавить докеризированный dev-frontend с проксированием на gateway.

Состав команды:
Александр Ермаков: Frontend/Backend
Дмитрий Карпов: Frontend
Эдуард Брага: Backend
Антон Лещев: Backend
