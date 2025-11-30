# API Documentation

## Authentication Endpoints

### Login

**Endpoint:** `POST /auth/login`

**Content-Type:** `application/json`

**Request Body:**
```json
{
  "email": "String",
  "password": "String"
}
```

**Response:**
```json
{
  "jwtToken": "String",
  "refreshToken": "String"
}
```

### Refresh Token

**Endpoint:** `POST /auth/refresh`

**Content-Type:** `application/json`

**Request Body:**
```json
{
  "refreshToken": "String"
}
```

**Response:**
```json
{
  "jwtToken": "String"
}
```

## Courses Endpoints

### Get Courses List

**Endpoint:** `GET /api/courses?page=Integer&size=Integer`

**Authorization:** `Bearer <jwtToken>`

**Response:**
```json
[
  {
    "courseId": "Integer",
    "courseName": "String",
    "difficulty": "String",
    "description": "String"
  }
]
```

### Get Course Details

**Endpoint:** `GET /api/courses/{courseId}`

**Authorization:** `Bearer <jwtToken>`

**Response:**
```json
{
  "courseId": "Integer",
  "title": "String",
  "description": "String",
  "modules": [
    {
      "moduleId": "Integer",
      "title": "String",
      "orderIndex": "Integer",
      "tasks": [
        {
          "taskId": "Integer",
          "title": "String",
          "type": "String",
          "duration": "String"
        },
        {
          "taskId": "Integer",
          "title": "String",
          "type": "String",
          "deadline": "String"
        }
      ]
    }
  ]
}
```

## Progress & Assignments

### Update Progress

**Endpoint:** `POST /api/progress`

**Authorization:** `Bearer <jwtToken>`

**Content-Type:** `application/json`

**Request Body:**
```json
{
  "taskId": "Integer",
  "status": "String",
  "score": "Integer",
  "completion_date": "String[Datetime]",
  "progress_id": "Integer",
  "student_id": "Integer",
}
```

**Response:**
```json
{
  "status": "Integer"
}
```

### Upload Assignment

**Endpoint:** `POST /assignments/upload`

**Authorization:** `Bearer <jwtToken>`

**Content-Type:** `multipart/form-data`

**Request Body:**
```json
{
  "studentId": "String",
  "taskId": "Integer",
  "file": "binary_file"
}
```

**Response:**
```json
{
  "status": "Integer",
  "assignmentID": "Integer"
}
```


