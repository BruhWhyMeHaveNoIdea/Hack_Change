-- Up migration: create schema for courses, modules, tasks, students, progress,
-- assignments, competencies and feedback_tickets
BEGIN;

CREATE TABLE courses (
    course_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    difficulty_level VARCHAR(255)
);

CREATE TABLE modules (
    module_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    course_id INTEGER NOT NULL REFERENCES courses(course_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    order_index INTEGER NOT NULL
);

CREATE TABLE tasks (
    task_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    module_id INTEGER NOT NULL REFERENCES modules(module_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    task_info TEXT,
    task_type VARCHAR(255) NOT NULL,
    points_value INTEGER DEFAULT 0
);

CREATE TABLE students (
    student_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    registration_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE progress (
    progress_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    student_id INTEGER NOT NULL REFERENCES students(student_id) ON DELETE CASCADE,
    task_id INTEGER NOT NULL REFERENCES tasks(task_id) ON DELETE CASCADE,
    status VARCHAR(255) NOT NULL,
    completion_date TIMESTAMPTZ,
    score REAL
);

CREATE TABLE assignments (
    assignment_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    student_id INTEGER NOT NULL REFERENCES students(student_id) ON DELETE CASCADE,
    task_id INTEGER NOT NULL REFERENCES tasks(task_id) ON DELETE CASCADE,
    upload_date TIMESTAMPTZ DEFAULT now(),
    file_path TEXT,
    status VARCHAR(50) DEFAULT 'Черновик'
);

CREATE TABLE competencies (
    competency_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    student_id INTEGER NOT NULL REFERENCES students(student_id) ON DELETE CASCADE,
    skill_name VARCHAR(100) NOT NULL,
    gap_level INTEGER CHECK (gap_level BETWEEN 1 AND 5) DEFAULT 1,
    last_updated TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE feedback_tickets (
    ticket_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    assignment_id INTEGER NOT NULL REFERENCES assignments(assignment_id) ON DELETE CASCADE,
    competency_id INTEGER NOT NULL REFERENCES competencies(competency_id) ON DELETE CASCADE,
    description TEXT,
    status VARCHAR(50) DEFAULT 'To Do',
    priority INTEGER DEFAULT 3 CHECK (priority >= 1 AND priority <= 3)
);

-- Indexes for faster lookups
CREATE INDEX idx_modules_course_id ON modules(course_id);
CREATE INDEX idx_tasks_module_id ON tasks(module_id);
CREATE INDEX idx_progress_student_id ON progress(student_id);
CREATE INDEX idx_progress_task_id ON progress(task_id);
CREATE INDEX idx_assignments_student_id ON assignments(student_id);
CREATE INDEX idx_assignments_task_id ON assignments(task_id);
CREATE INDEX idx_competencies_student_id ON competencies(student_id);
CREATE INDEX idx_feedback_assignment_id ON feedback_tickets(assignment_id);

COMMIT;
