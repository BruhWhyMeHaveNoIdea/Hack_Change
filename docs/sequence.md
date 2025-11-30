##  Пользовательский сценарий: Трек «Студент»

Решение реализует полный путь студента в соответствии с ТЗ:
- просмотр учебных материалов,
- загрузка домашнего задания (до 100 МБ),
- получение комментария и оценки от преподавателя.

### Последовательность взаимодействия компонентов

```mermaid
sequenceDiagram
    participant Frontend as Frontend (React)
    participant Gateway as Gateway & Auth (Python)
    participant ContentSvc as Content Service
    participant AssignmentSvc as Assignment Service
    participant DB as PostgreSQL

    %% === Авторизация ===
    Frontend->>Gateway: POST /auth/login {email, password}
    Gateway->>DB: SELECT * FROM students WHERE email = ?
    DB-->>Gateway: student data
    Gateway->>Gateway: generate JWT + refreshToken
    Gateway-->>Frontend: 200 OK {jwtToken, refreshToken}

    %% === Обновление токена (опционально) ===
    Frontend->>Gateway: POST /auth/refresh {refreshToken}
    Gateway->>Gateway: validate + issue new JWT
    Gateway-->>Frontend: 200 OK {jwtToken}

    %% === Получение списка курсов ===
    Frontend->>Gateway: GET /api/courses?page=0&size=10
    Gateway->>Gateway: validate JWT
    Gateway->>ContentSvc: GET /courses?page=0&size=10
    ContentSvc->>DB: SELECT course_id, title, difficulty_level FROM courses LIMIT 10 OFFSET 0
    DB-->>ContentSvc: course list
    ContentSvc-->>Gateway: [courseId, courseName, difficulty]
    Gateway-->>Frontend: 200 [courses]

    %% === Получение деталей курса + модулей + заданий ===
    Frontend->>Gateway: GET /api/courses/{courseId}
    Gateway->>Gateway: validate JWT
    Gateway->>ContentSvc: GET /courses/{courseId}
    ContentSvc->>DB: SELECT c.*, m.module_id, m.title AS module_title, m.order_index<br/>FROM courses c<br/>JOIN modules m ON c.course_id = m.course_id<br/>JOIN tasks t ON m.module_id = t.module_id<br/>WHERE c.course_id = ?
    DB-->>ContentSvc: course + modules + tasks
    ContentSvc-->>Gateway: {courseId, title, description, modules: [{moduleId, title, tasks: [...]}]}
    Gateway-->>Frontend: 200 course with full structure

    %% === Загрузка домашнего задания (тип Assignment) ===
    Frontend->>Gateway: POST /api/assignments/upload {taskId, file}
    Gateway->>Gateway: validate JWT → extract studentId
    Gateway->>AssignmentSvc: POST /upload {studentId, taskId, file}
    AssignmentSvc->>DB: INSERT INTO assignments (student_id, task_id, file_path, status)<br/>VALUES (?, ?, ?, 'На проверке')
    DB-->>AssignmentSvc: assignment_id
    AssignmentSvc-->>Gateway: 201 {assignmentId}
    Gateway-->>Frontend: 201 OK

    %% === Обновление прогресса (например, после выполнения Quiz/Video) ===
    Frontend->>Gateway: POST /api/progress {taskId, status: "Completed", score: 90}
    Gateway->>Gateway: validate JWT → studentId
    Gateway->>ContentSvc: POST /progress {studentId, taskId, status, score}
    ContentSvc->>DB: INSERT INTO progress (student_id, task_id, status, score, completion_date)<br/>VALUES (?, ?, ?, ?, NOW())
    DB-->>ContentSvc: OK
    ContentSvc-->>Gateway: 201 OK
    Gateway-->>Frontend: 201 OK

    %% === Получение обратной связи по заданию (через assignmentId) ===
    Note over Frontend,Gateway: Через некоторое время (после проверки преподавателем)

    Frontend->>Gateway: GET /api/assignments/{assignmentId}/feedback
    Gateway->>Gateway: validate JWT
    Gateway->>AssignmentSvc: GET /{assignmentId}/feedback
    AssignmentSvc->>DB: SELECT ft.description, ft.priority, ft.status, c.skill_name<br/>FROM feedback_tickets ft<br/>JOIN competencies c ON ft.competency_id = c.competency_id<br/>WHERE ft.assignment_id = ?
    DB-->>AssignmentSvc: feedback tickets + competencies
    AssignmentSvc-->>Gateway: [{description, skill, priority, status}]
    Gateway-->>Frontend: 200 [feedback]

    %% === Просмотр общей статистики (опционально — шкала прогресса) ===
    Frontend->>Gateway: GET /api/progress/summary?courseId=123
    Gateway->>Gateway: validate JWT → studentId
    Gateway->>ContentSvc: GET /progress/summary?studentId=...&courseId=123
    ContentSvc->>DB: SELECT COUNT(t.task_id) AS total,<br/>COUNT(p.task_id) AS completed<br/>FROM tasks t<br/>LEFT JOIN progress p ON t.task_id = p.task_id AND p.student_id = ?<br/>JOIN modules m ON t.module_id = m.module_id<br/>WHERE m.course_id = ?
    DB-->>ContentSvc: {total, completed}
    ContentSvc-->>Gateway: {total: 25, completed: 18}
    Gateway-->>Frontend: 200 {progress: "72%"}
```