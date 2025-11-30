<?php

function load_env($path = __DIR__ . '/.env') {
    if (!file_exists($path))
        return;

    $lines = file($path, file_ignore_empty_lines | file_skip_new_lines);

    foreach ($lines as $line) {
        // Пропускаем комментарии
        if(str_starts_with(trim($line), '#'))
            continue;

        // Разбиваем строку на ключ и значение
        if (strpos($line, '=') !== false) {
            list($name, $value) = explode('=', $line, 2);
            $name = trim($name);
            $value = trim($value, " \t\n\r\0\x0B\"'");

            // Устанавливаем переменную окружения
            if (!empty($name)) {
                putenv("{$name}={$value}");
                $_ENV[$name] = $value;
                $_SERVER[$name] = $value;
            }
        }
    }
}

load_env();

function env(string $key, $default = null) {
    // Сначала ищем в $_ENV, затем в getenv()
    $value = $_ENV[$key] ?? getenv($key);

    // Если значение - 'null', возвращаем null
    if ($value === 'null')
        return null;

    // Если переменная не найдена, возвращаем значение по умолчанию
    if ($value === false)
        return $default;

    return $value;
}

function getDbConnection(): PDO {
    // Получаем данные из переменных окружения
    $db_host = env('POSTGRES_HOST', 'localhost');
    $db_port = env('POSTGRES_PORT', '5432');
    $db_name = env('POSTGRES_DB', 'lms');
    $db_user = env('POSTGRES_USER', 'postgres');
    $db_pass = env('POSTGRES_PASSWORD', 'postgres');

    if (!$db_name || !$db_user || !$db_pass) {
        throw new Exception("Missing required DB configuration in .env");
    }

    $dsn = "pgsql:host=" . $db_host . ";port=" . $db_port . ";dbname=" . $db_name;
    
    $options = [
        PDO::ATTR_ERRMODE            => PDO::ERRMODE_EXCEPTION,
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
        PDO::ATTR_EMULATE_PREPARES   => false,
    ];

    try {
        $pdo = new PDO($dsn, $db_user, $db_pass, $options);
        return $pdo;
    } catch (PDOException $e) {
        error_log("Ошибка подключения к БД: " . $e->getMessage());
        http_response_code(500);
        exit(json_encode(["error" => "Database connection failed"]));
    }
}