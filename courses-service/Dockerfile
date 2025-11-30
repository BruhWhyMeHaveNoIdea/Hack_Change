# Используем официальный образ PHP с PHP-FPM на базе Alpine
FROM php:8.2-fpm-alpine

# Установка системных зависимостей и PHP-расширений для работы с PostgreSQL и кеширования
RUN apk update && apk add --no-cache \
    # PostgreSQL клиент и библиотеки разработки (libpq-dev нужен для сборки расширений)
    postgresql-client \
    libpq-dev \
    \
    # Добавляем инструменты сборки ($PHPIZE_DEPS) временно
    $PHPIZE_DEPS \
    \
    # Устанавливаем расширения PHP: pdo_pgsql, pgsql (PostgreSQL) и opcache (производительность)
    && docker-php-ext-install pdo_pgsql pgsql opcache \
    \
    # Удаляем инструменты сборки сразу после установки расширений для минимизации размера образа
    && apk del $PHPIZE_DEPS \
    && rm -rf /var/cache/apk/*

# Устанавливаем рабочую директорию
WORKDIR /var/www/html

# Копируем весь ваш код из текущей папки сборки в контейнер
# (Предполагается, что Dockerfile находится в корневой папке вашего приложения)
COPY . /var/www/html

# Устанавливаем правильные права доступа для пользователя www-data, под которым работает FPM
RUN chown -R www-data:www-data /var/www/html

# PHP-FPM по умолчанию слушает на порту 9000
EXPOSE 9000

# Команда, которая запускает PHP-FPM
CMD ["php-fpm"]