-- Seed data migration
BEGIN;

-- 1. Вставляем курсы (без изменений)
INSERT INTO courses (title, description, difficulty_level) VALUES
('Основы SQL и реляционных баз данных', 'Изучите основы SQL и работу с реляционными базами данных.', 'Начальный'),
('Алгоритмы и структуры данных', 'Познакомьтесь с основными алгоритмами и структурами данных.', 'Средний'),
('Микросервисная архитектура', 'Освойте принципы проектирования и развертывания микросервисов.', 'Продвинутый'),
('Введение в DevOps', 'Научитесь автоматизировать процессы разработки и развертывания.', 'Начальный'),
('Безопасность веб-приложений', 'Изучите основы защиты веб-приложений от угроз.', 'Средний'),
('Распределённые системы', 'Понимание принципов работы распределённых систем и их масштабирования.', 'Продвинутый');

-- 2. Создаём модуль для первого курса (course_id = 1)
INSERT INTO modules (course_id, title, order_index) VALUES
(1, 'Основы SQL', 1);

-- 3. Добавляем 5 задач в модуль (module_id = 1)
INSERT INTO tasks (module_id, title, task_info, task_type, points_value) VALUES
(1, 'Создание таблицы', 'Напишите SQL-запрос для создания таблицы users с полями: id (INT, PRIMARY KEY), name (VARCHAR), email (VARCHAR).', 'code', 10),
(1, 'Выборка данных', 'Напишите запрос для выборки всех пользователей, у которых email заканчивается на "@example.com".', 'code', 10),
(1, 'Ошибочный запрос', 'Исправьте ошибку в запросе: SELECT * FROM users WHER email = ''test@example.com'';', 'code', 10),
(1, 'Типы данных в SQL', 'Перечислите 5 основных типов данных в PostgreSQL и кратко опишите каждый.', 'text', 10),
(1, 'Первичный и внешний ключ', 'Объясните разницу между PRIMARY KEY и FOREIGN KEY. Приведите пример с двумя таблицами.', 'text', 10);

-- 4. Создаём студента
INSERT INTO students (first_name, last_name, email, password) VALUES
('Иван', 'Иванов', 'ivan@example.com', '$2a$10$c83aASYEZCKupTWCLSkVxO0d.SCFObdtg1s8iXXcK4IaR2fpUav0S');

-- Предположим: student_id = 1, task_id от 1 до 5

-- 5. Прогресс по первым трём задачам
INSERT INTO progress (student_id, task_id, status, completion_date, score) VALUES
(1, 1, 'completed', now(), 10.0),
(1, 2, 'completed', now(), 9.5),
(1, 3, 'needs_revision', now(), 4.0);  -- ошибка, нужно переделать

-- 6. Отправки (assignments) для первых трёх задач
INSERT INTO assignments (student_id, task_id, file_path, status) VALUES
(1, 1, '/uploads/ivan/task1.sql', 'Отправлено'),
(1, 2, '/uploads/ivan/task2.sql', 'Отправлено'),
(1, 3, '/uploads/ivan/task3.sql', 'Отправлено');

-- 7. Компетенция студента
INSERT INTO competencies (student_id, skill_name, gap_level) VALUES
(1, 'SQL', 2);

-- Получаем assignment_id и competency_id для связи (предположим: assignment_id = 1,2,3; competency_id = 1)

-- 8. Фидбэки (feedback_tickets)
-- Для первых двух — проверено, для третьего — требует доработки
INSERT INTO feedback_tickets (assignment_id, competency_id, description, status, priority) VALUES
(1, 1, 'Отличная работа! Запрос корректный и соответствует ТЗ.', 'Done', 1),
(2, 1, 'Хорошо, но можно использовать ILIKE для case-insensitive поиска.', 'Done', 2),
(3, 1, 'В запросе опечатка: WHER вместо WHERE. Исправьте и отправьте повторно.', 'To Do', 1);

COMMIT;