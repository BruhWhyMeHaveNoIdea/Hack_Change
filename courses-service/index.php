<?php
// Описание: Главный файл API для Content Service. 
// Служит центральной точкой для обработки всех запросов, связанных с контентом и прогрессом пользователей (CRUD-операции с курсами, модулями, заданиями и отслеживанием прогресса).

// Устанавливаем критические HTTP-заголовки
header('Content-Type: application/json'); // Все ответы будут в формате JSON

// Конфигурация CORS (Внимание: '*' подходит только для разработки)
header('Access-Control-Allow-Origin: *'); 
header('Access-Control-Allow-Methods: GET, POST, OPTIONS');
header('Access-Control-Allow-Headers: Content-Type, Authorization');

// Включаем конфигурацию базы данных (файл db.php)
require_once 'db.php';

// --- ГЛОБАЛЬНЫЕ КОНСТАНТЫ ---
// Допустимые статусы для отслеживания прогресса пользователя
const ALLOWED_STATUSES = ['Started', 'Completed', 'Failed', 'InProgress'];

// Получаем метод HTTP-запроса (GET, POST и т.д.)
$requestMethod = $_SERVER['REQUEST_METHOD'];
// Извлекаем путь URI без домена и параметров (например, /courses/0)
$requestPath = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);
// Разбиваем путь на сегменты (например, ['courses', '0'])
$pathSegments = explode('/', trim($requestPath, '/'));

// --- ВСПОМОГАТЕЛЬНАЯ ФУНКЦИЯ ДЛЯ ОТВЕТА ---

/**
 * Отправляет JSON-ответ и завершает выполнение скрипта.
 * @param mixed $data Данные для кодирования в JSON.
 * @param int $statusCode HTTP-код ответа (по умолчанию 200).
 * @return void
 */
function sendJsonResponse($data, $statusCode = 200) {
    http_response_code($statusCode);
    // Используем флаг JSON_UNESCAPED_UNICODE для корректного отображения кириллицы
    echo json_encode($data, JSON_UNESCAPED_UNICODE);
    exit;
}

// --- ПРЕДВАРИТЕЛЬНАЯ ОБРАБОТКА (CORS Pre-flight) ---
if ($requestMethod === 'OPTIONS') {
    // Для OPTIONS-запросов (CORS) достаточно вернуть 204 No Content
    sendJsonResponse(null, 204);
}

// --- РОУТИНГ: ИДЕНТИФИКАЦИЯ РЕСУРСА ---
// Извлекаем первый сегмент как основной ресурс (например, 'courses' или 'progress')
$mainResource = array_shift($pathSegments); 

if (empty($mainResource)) {
    // Если в пути нет ресурса (запрос к корню), возвращаем 404
    sendJsonResponse(['error' => 'Not Found'], 404);
}

// --- ПОДКЛЮЧЕНИЕ К БАЗЕ ДАННЫХ ---
try {
    // $pdo - это объект PDO для взаимодействия с БД
    $pdo = getDbConnection(); 
} catch (Exception $e) {
    // В случае критической ошибки подключения
    sendJsonResponse(['error' => 'Database connection failed: ' . $e->getMessage()], 500);
}


// --- ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ДЛЯ ПРОВЕРОК В БД ---

/**
 * Проверяет существование студента по ID. В случае отсутствия отправляет 404.
 * @param PDO $pdo
 * @param string $studentId
 * @return void
 */
function checkStudentExists(PDO $pdo, string $studentId): void {
    $sql = "SELECT student_id FROM students WHERE student_id = :studentId";
    $stmt = $pdo->prepare($sql);
    $stmt->bindValue(':studentId', $studentId, PDO::PARAM_STR);
    $stmt->execute();
    if (!$stmt->fetch()) {
        sendJsonResponse(['error' => "Student with ID {$studentId} not found."], 404);
    }
}

/**
 * Проверяет существование курса по ID. В случае отсутствия отправляет 404.
 * @param PDO $pdo
 * @param int $courseId
 * @return void
 */
function checkCourseExists(PDO $pdo, int $courseId): void {
    $sql = "SELECT course_id FROM courses WHERE course_id = :courseId";
    $stmt = $pdo->prepare($sql);
    $stmt->bindValue(':courseId', $courseId, PDO::PARAM_INT);
    $stmt->execute();

    if (!$stmt->fetch()) {
        sendJsonResponse(['error' => "Course with ID {$courseId} not found."], 404);
    }
}

/**
 * Проверяет существование задания по ID. В случае отсутствия отправляет 404.
 * @param PDO $pdo
 * @param int $taskId
 * @return void
 */
function checkTaskExists(PDO $pdo, int $taskId): void {
    $sql = "SELECT task_id FROM tasks WHERE task_id = :taskId";
    $stmt = $pdo->prepare($sql);
    $stmt->bindValue(':taskId', $taskId, PDO::PARAM_INT);
    $stmt->execute();

    if (!$stmt->fetch()) {
        sendJsonResponse(['error' => "Task with ID {$taskId} not found."], 404);
    }
}

/**
 * Проверяет допустимость данных прогресса (статус и оценка).
 * @param array $data Массив данных из тела POST-запроса.
 * @return void
 */
function validateProgressData(array $data): void {
    $status = $data['status'];
    $score = $data['score'] ?? null;
    
    // 1. Проверка статуса: должен быть одним из разрешенных
    if (!in_array($status, ALLOWED_STATUSES)) {
        $allowed = implode(', ', ALLOWED_STATUSES);
        sendJsonResponse(['error' => "Invalid status value: '{$status}'. Must be one of: {$allowed}"], 400);
    }

    // 2. Проверка оценки: если есть, должна быть неотрицательным целым числом
    if ($score !== null) {
        if (!is_int($score) || $score < 0) {
            sendJsonResponse(['error' => 'Score must be a non-negative integer.'], 400);
        }
    }
}

// --- ГЛАВНЫЙ БЛОК РОУТИНГА ---

try {
    switch ($mainResource) {
        
        // 1. Обработка всех запросов, начинающихся с /courses
        case 'courses':
            
            // 1.1 GET /courses (Получение списка курсов с пагинацией)
            if ($requestMethod === 'GET' && count($pathSegments) === 0) {
                // Извлекаем параметры пагинации из QUERY STRING
                $page = max(1, (int) ($_GET['page'] ?? 1));
                $pageSize = max(1, (int) ($_GET['size'] ?? 10));
                
                $offset = ($page - 1) * $pageSize;

                $sql = "SELECT course_id AS id, title, description, difficulty_level AS difficulty 
                        FROM courses 
                        ORDER BY course_id 
                        LIMIT :limit OFFSET :offset";

                $stmt = $pdo->prepare($sql);
                $stmt->bindValue(':limit', $pageSize, PDO::PARAM_INT);
                $stmt->bindValue(':offset', $offset, PDO::PARAM_INT);
                $stmt->execute();
                
                $coursesList = $stmt->fetchAll(PDO::FETCH_ASSOC);

                sendJsonResponse($coursesList);
            } 
            
            // 1.2 GET /courses/{courseId} (Получение полной структуры одного курса)
            elseif ($requestMethod === 'GET' && count($pathSegments) === 1) {
                $courseIdSegment = $pathSegments[0]; 
                
                // Проверка: ID должен быть числом
                if (!is_numeric($courseIdSegment)) {
                    sendJsonResponse(['error' => "Invalid Course ID format: '{$courseIdSegment}'"], 404);
                }

                $courseId = (int) $courseIdSegment;

                // Проверка существования курса (вызовет 404, если не найден)
                checkCourseExists($pdo, $courseId);

                // --- 1. Получение основной информации о курсе ---
                $courseSql = "SELECT course_id AS id, title, description, difficulty_level AS difficulty 
                              FROM courses 
                              WHERE course_id = :courseId";
                $courseStmt = $pdo->prepare($courseSql);
                $courseStmt->bindValue(':courseId', $courseId, PDO::PARAM_INT);
                $courseStmt->execute();
                $courseData = $courseStmt->fetch(PDO::FETCH_ASSOC);

                // --- 2. Получение модулей и заданий вложенной структурой ---
                $structureSql = "
                    SELECT 
                        m.module_id, 
                        m.title AS module_title, 
                        m.order_index, 
                        t.task_id, 
                        t.title AS task_title, 
                        t.task_type,
                        t.points_value
                    FROM modules m
                    LEFT JOIN tasks t ON m.module_id = t.module_id
                    WHERE m.course_id = :courseId
                    ORDER BY m.order_index, t.task_id
                ";

                $structureStmt = $pdo->prepare($structureSql);
                $structureStmt->bindValue(':courseId', $courseId, PDO::PARAM_INT);
                $structureStmt->execute();
                $rows = $structureStmt->fetchAll(PDO::FETCH_ASSOC);

                $modulesTree = [];
                foreach ($rows as $row) {
                    $moduleId = $row['module_id'];
                    if (!isset($modulesTree[$moduleId])) {
                        $modulesTree[$moduleId] = [
                            'module_id' => $moduleId,
                            'title' => $row['module_title'],
                            'order_index' => $row['order_index'],
                            'tasks' => []
                        ];
                    }
                    // Добавляем задание, если оно не NULL
                    if ($row['task_id'] !== null) {
                        $modulesTree[$moduleId]['tasks'][] = [
                            'task_id' => $row['task_id'],
                            'title' => $row['task_title'],
                            'task_type' => $row['task_type'],
                            'points_value' => $row['points_value'],
                        ];
                    }
                }
                
                $courseData['modules'] = array_values($modulesTree);

                sendJsonResponse($courseData);

            } else {
                // Неподдерживаемый метод или неверная структура URL
                sendJsonResponse(['error' => 'Method Not Allowed or Invalid Course URL Structure'], 405);
            }
            break;


        // 2. Обработка всех запросов, начинающихся с /progress
        case 'progress':
            
            $progressSegment = array_shift($pathSegments);

            // 2.1 GET /progress/summary?studentId=...&courseId=... (Получение сводного прогресса)
            if ($requestMethod === 'GET' && strtolower(trim($progressSegment)) === 'summary' && count($pathSegments) === 0) {
                
                $studentId = $_GET['studentId'] ?? null; 
                $courseId = $_GET['courseId'] ?? null; 

                if ($studentId === null || $courseId === null) {
                    sendJsonResponse(['error' => 'Missing required query parameters: studentId and courseId'], 400);
                }
                
                // Проверки существования студента и курса
                checkStudentExists($pdo, $studentId); 
                checkCourseExists($pdo, (int)$courseId);
                
                // SQL-запрос для подсчета общего и завершенного числа заданий
                $sql = "
                    SELECT 
                        COUNT(t.task_id) AS total, 
                        COUNT(p.task_id) FILTER (WHERE p.status = 'Completed') AS completed
                    FROM tasks t
                    JOIN modules m ON t.module_id = m.module_id
                    LEFT JOIN progress p ON t.task_id = p.task_id AND p.student_id = :studentId
                    WHERE m.course_id = :courseId
                ";

                $stmt = $pdo->prepare($sql);
                $stmt->bindValue(':studentId', $studentId, PDO::PARAM_STR); 
                $stmt->bindValue(':courseId', (int) $courseId, PDO::PARAM_INT);
                $stmt->execute();
                
                $summaryResult = $stmt->fetch(PDO::FETCH_ASSOC);

                sendJsonResponse([
                    'total' => (int) ($summaryResult['total'] ?? 0),
                    'completed' => (int) ($summaryResult['completed'] ?? 0)
                ]);
            
            // 2.2 POST /progress (Сохранение/обновление прогресса задания)
            } elseif ($requestMethod === 'POST' && $progressSegment === null) {
                
                $jsonPayload = file_get_contents('php://input');
                $progressData = json_decode($jsonPayload, true);

                if (json_last_error() !== JSON_ERROR_NONE) {
                    sendJsonResponse(['error' => 'Invalid JSON input'], 400);
                }

                // Проверка обязательных полей
                if (empty($progressData['studentId']) || empty($progressData['taskId']) || empty($progressData['status'])) {
                    sendJsonResponse(['error' => 'Missing required fields: studentId, taskId, and status'], 400);
                }

                $studentId = $progressData['studentId'];
                $taskId = (int) $progressData['taskId'];
                $status = $progressData['status'];
                // Оценка может быть опущена (null)
                $score = $progressData['score'] ?? null; 

                // Валидация данных (проверка статуса и оценки)
                validateProgressData($progressData);
                
                // Проверки существования студента и задания
                checkStudentExists($pdo, $studentId); 
                checkTaskExists($pdo, $taskId); 

                // PostgreSQL UPSERT: Вставить или Обновить. Гарантирует атомарность.
                $sql = "
                    INSERT INTO progress (student_id, task_id, status, score, last_updated_at)
                    VALUES (:studentId, :taskId, :status, :score, NOW())
                    ON CONFLICT (student_id, task_id) DO UPDATE
                    SET status = EXCLUDED.status, 
                        score = EXCLUDED.score, 
                        last_updated_at = NOW()
                    RETURNING progress_id, student_id, task_id, status, score, last_updated_at
                ";

                $stmt = $pdo->prepare($sql);
                $stmt->bindValue(':studentId', $studentId, PDO::PARAM_STR);
                $stmt->bindValue(':taskId', $taskId, PDO::PARAM_INT);
                $stmt->bindValue(':status', $status, PDO::PARAM_STR);
                // Корректная привязка NULL к PDO
                $stmt->bindValue(':score', $score, $score === null ? PDO::PARAM_NULL : PDO::PARAM_INT);
                
                $stmt->execute();
                
                $updatedProgressRecord = $stmt->fetch(PDO::FETCH_ASSOC);

                sendJsonResponse($updatedProgressRecord, 200);

            } else {
                // Если запрос не POST /progress и не GET /progress/summary
                sendJsonResponse(['error' => 'Not Found'], 404);
            }
            break;

        default:
            // Обработка любого другого несуществующего ресурса
            sendJsonResponse(['error' => 'Not Found'], 404);
            break;
    }

} catch (PDOException $e) {
    // Обработка ошибок, связанных с базой данных
    error_log("Ошибка БД: " . $e->getMessage());
    sendJsonResponse(['error' => 'Internal Database Error'], 500); 
} catch (Exception $e) {
    // Обработка любых других неожиданных внутренних ошибок
    error_log("Ошибка: " . $e->getMessage());
    sendJsonResponse(['error' => 'Internal Server Error'], 500);
}