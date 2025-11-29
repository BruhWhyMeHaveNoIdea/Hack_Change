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

    Frontend->>Gateway: POST /auth/login {email, password}
    Gateway->>DB: SELECT user WHERE email = ?
    DB-->>Gateway: user data
    Gateway->>Gateway: generate JWT and refreshToken
    Gateway-->>Frontend: 200 OK {jwtToken, refreshToken}

    Frontend->>Gateway: POST /auth/refresh {refreshToken}
    Gateway->>DB: SELECT user WHERE email = ?
    DB-->>Gateway: user data
    Gateway->>Gateway: generate JWT
    Gateway-->>Frontend: 200 OK {jwtToken}
```