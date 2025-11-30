# Используем CLI-образ (а не FPM!), потому что нам нужен HTTP-сервер
FROM php:8.2-cli-alpine

# Установка расширений PostgreSQL
RUN apk update && apk add --no-cache \
    postgresql-client libpq-dev \
    && docker-php-ext-install pdo_pgsql pgsql \
    && rm -rf /var/cache/apk/*

# Рабочая директория
WORKDIR /var/www/html

# Копируем код
COPY . .

# Даём права (опционально, но хорошо для безопасности)
RUN chown -R www-data:www-data /var/www/html

# Переключаемся на непривилегированного пользователя
USER www-data

# Встроенный сервер будет слушать порт 9000
EXPOSE 9000

# Запуск сервера на 0.0.0.0:9000
CMD ["php", "-S", "0.0.0.0:9000", "index.php"]