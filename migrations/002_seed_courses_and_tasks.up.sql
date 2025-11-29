-- Up migration: Insert sample courses, modules, and tasks based on UI design
BEGIN;

-- Вставляем курсы
INSERT INTO courses (title, description, difficulty_level) VALUES
('Основы SQL и реляционных баз данных', 'Изучите основы SQL и работу с реляционными базами данных.', 'Начальный'),
('Алгоритмы и структуры данных', 'Познакомьтесь с основными алгоритмами и структурами данных.', 'Средний'),
('Микросервисная архитектура', 'Освойте принципы проектирования и развертывания микросервисов.', 'Продвинутый'),
('Введение в DevOps', 'Научитесь автоматизировать процессы разработки и развертывания.', 'Начальный'),
('Безопасность веб-приложений', 'Изучите основы защиты веб-приложений от угроз.', 'Средний'),
('Распределённые системы', 'Понимание принципов работы распределённых систем и их масштабирования.', 'Продвинутый');

-- Вставляем модули для каждого курса (по одному модулю на курс)
-- Используем CTE для получения course_id и последующего вставления модулей

WITH inserted_courses AS (
    SELECT course_id, title FROM courses WHERE title IN (
        'Основы SQL и реляционных баз данных',
        'Алгоритмы и структуры данных',
        'Микросервисная архитектура',
        'Введение в DevOps',
        'Безопасность веб-приложений',
        'Распределённые системы'
    )
)
INSERT INTO modules (course_id, title, order_index)
SELECT course_id, 'Введение в курс: ' || title, 1
FROM inserted_courses;

-- Вставляем по одному простому заданию в каждый модуль
-- Используем CTE для получения module_id и последующего вставления задач
-- Для каждого модуля вставим несколько типовых заданий (Video, Quiz, Assignment)
WITH inserted_modules AS (
    SELECT m.module_id, c.title AS course_title
    FROM modules m
    JOIN courses c ON m.course_id = c.course_id
    WHERE c.title IN (
        'Основы SQL и реляционных баз данных',
        'Алгоритмы и структуры данных',
        'Микросервисная архитектура',
        'Введение в DevOps',
        'Безопасность веб-приложений',
        'Распределённые системы'
    )
)
INSERT INTO tasks (module_id, title, task_type, task_info, points_value)
SELECT module_id,
       'Видео: ' || course_title,
       'Video',
       'Короткое вводное видео по теме: ' || course_title,
       5
FROM inserted_modules
UNION ALL
SELECT module_id,
       'Тест: ' || course_title,
       'Quiz',
       'Небольшой тест для самопроверки по теме: ' || course_title,
       10
FROM inserted_modules
UNION ALL
SELECT module_id,
       'Практическое задание: ' || course_title,
       'Assignment',
       'Практическое задание с загрузкой решения для проверки преподавателем',
       20
FROM inserted_modules;

COMMIT;
