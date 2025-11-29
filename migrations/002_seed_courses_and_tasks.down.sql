-- Down migration: remove seeded courses, modules and tasks inserted by 002_seed_courses_and_tasks.up.sql
BEGIN;

-- Delete tasks that belong to the modules we created for these courses
DELETE FROM tasks
WHERE module_id IN (
    SELECT m.module_id
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
);

-- Delete modules that were inserted for the seeded courses
DELETE FROM modules
WHERE course_id IN (
    SELECT course_id FROM courses WHERE title IN (
        'Основы SQL и реляционных баз данных',
        'Алгоритмы и структуры данных',
        'Микросервисная архитектура',
        'Введение в DevOps',
        'Безопасность веб-приложений',
        'Распределённые системы'
    )
);

-- Delete the courses themselves
DELETE FROM courses
WHERE title IN (
    'Основы SQL и реляционных баз данных',
    'Алгоритмы и структуры данных',
    'Микросервисная архитектура',
    'Введение в DevOps',
    'Безопасность веб-приложений',
    'Распределённые системы'
);

COMMIT;
