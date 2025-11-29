##  Пользовательский сценарий: Трек «Студент»

Решение реализует полный путь студента в соответствии с ТЗ:
- просмотр учебных материалов,
- загрузка домашнего задания (до 100 МБ),
- получение комментария и оценки от преподавателя.

### Последовательность взаимодействия компонентов

```mermaid
sequenceDiagram
    participant Frontend as Frontend (React)
    participant Gateway as Gateway & Auth (Java)
    participant ContentSvc as Content Service
    participant AssignmentSvc as Assignment Service
    participant DB as PostgreSQL

    Frontend->>Gateway: POST /auth/login {email, password}
    Gateway->>DB: SELECT user WHERE email = ?
    DB-->>Gateway: user data
    Gateway->>Gateway: generate JWT
    Gateway-->>Frontend: 200 OK {token}

    Frontend->>Gateway: GET /api/courses
    Gateway->>Gateway: validate JWT
    Gateway->>ContentSvc: GET /courses
    ContentSvc->>DB: SELECT courses, materials
    DB-->>ContentSvc: course list
    ContentSvc-->>Gateway: course data
    Gateway-->>Frontend: 200 [courses]

    Frontend->>Gateway: POST /api/assignments/upload {file, assignmentId}
    Gateway->>Gateway: validate JWT
    Gateway->>AssignmentSvc: POST /upload {file, userId, assignmentId}
    AssignmentSvc->>DB: INSERT assignment_version
    DB-->>AssignmentSvc: OK
    AssignmentSvc-->>Gateway: 201 Created
    Gateway-->>Frontend: 201 OK

    Note over Frontend,Gateway: (через некоторое время)

    Frontend->>Gateway: GET /api/assignments/{id}/feedback
    Gateway->>Gateway: validate JWT
    Gateway->>AssignmentSvc: GET /{id}/feedback
    AssignmentSvc->>DB: SELECT feedback, grade, comments
    DB-->>AssignmentSvc: feedback data
    AssignmentSvc-->>Gateway: feedback
    Gateway-->>Frontend: 200 {grade, comment}
```