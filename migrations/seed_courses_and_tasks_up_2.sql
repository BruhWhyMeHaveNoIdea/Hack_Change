/*
 Navicat Premium Dump SQL

 Source Server         : Postgres Local
 Source Server Type    : PostgreSQL
 Source Server Version : 180001 (180001)
 Source Host           : localhost:5432
 Source Catalog        : lms
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 180001 (180001)
 File Encoding         : 65001

 Date: 30/11/2025 03:22:58
*/


-- ----------------------------
-- Table structure for assigments
-- ----------------------------
DROP TABLE IF EXISTS "public"."assigments";
CREATE TABLE "public"."assigments" (
  "assignment_id" int4 NOT NULL,
  "student_id" int4,
  "task_id" int4,
  "upload_date" timestamp(6),
  "file_path" text COLLATE "pg_catalog"."default",
  "status" varchar(50) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."assigments" OWNER TO "postgres";
COMMENT ON COLUMN "public"."assigments"."assignment_id" IS 'Уникальный ID загруженного файла задания';
COMMENT ON COLUMN "public"."assigments"."student_id" IS 'Ссылка на студента, который сдал работу';
COMMENT ON COLUMN "public"."assigments"."task_id" IS 'Задание, к которому относится сданная работа';
COMMENT ON COLUMN "public"."assigments"."upload_date" IS 'Дата и время загрузки файла';
COMMENT ON COLUMN "public"."assigments"."file_path" IS 'Путь к сохраненному файлу на сервере (для S3)';
COMMENT ON COLUMN "public"."assigments"."status" IS 'Статус проверки: Черновик, На проверке, Проверено';

-- ----------------------------
-- Records of assigments
-- ----------------------------
BEGIN;
INSERT INTO "public"."assigments" ("assignment_id", "student_id", "task_id", "upload_date", "file_path", "status") VALUES (0, 0, 0, '2025-11-30 03:17:03.966469', '/uploads/ivan/task1.sql', 'На проверке');
INSERT INTO "public"."assigments" ("assignment_id", "student_id", "task_id", "upload_date", "file_path", "status") VALUES (1, 0, 1, '2025-11-30 03:17:03.966469', '/uploads/ivan/task2.sql', 'На проверке');
INSERT INTO "public"."assigments" ("assignment_id", "student_id", "task_id", "upload_date", "file_path", "status") VALUES (2, 0, 2, '2025-11-30 03:17:03.966469', '/uploads/ivan/task3.sql', 'На проверке');
COMMIT;

-- ----------------------------
-- Table structure for competencies
-- ----------------------------
DROP TABLE IF EXISTS "public"."competencies";
CREATE TABLE "public"."competencies" (
  "competency_id" int4 NOT NULL,
  "student_id" int4,
  "skill_name" varchar(100) COLLATE "pg_catalog"."default",
  "gap_level" int4,
  "last_updated" timestamp(6)
)
;
ALTER TABLE "public"."competencies" OWNER TO "postgres";
COMMENT ON COLUMN "public"."competencies"."competency_id" IS 'ID для отслеживания конкретного навыка/темы';
COMMENT ON COLUMN "public"."competencies"."student_id" IS 'Ссылка на студента';
COMMENT ON COLUMN "public"."competencies"."skill_name" IS 'Название навыка, в котором выявлен пробел';
COMMENT ON COLUMN "public"."competencies"."gap_level" IS 'Уровень дефицита (от 1 — низкий, до 5 — высокий)';
COMMENT ON COLUMN "public"."competencies"."last_updated" IS 'Дата, когда компетенция была оценена или обновлена';

-- ----------------------------
-- Records of competencies
-- ----------------------------
BEGIN;
INSERT INTO "public"."competencies" ("competency_id", "student_id", "skill_name", "gap_level", "last_updated") VALUES (0, 0, 'SQL', 2, '2025-11-30 03:18:17.414825');
COMMIT;

-- ----------------------------
-- Table structure for courses
-- ----------------------------
DROP TABLE IF EXISTS "public"."courses";
CREATE TABLE "public"."courses" (
  "course_id" int4 NOT NULL,
  "title" varchar(255) COLLATE "pg_catalog"."default",
  "description" text COLLATE "pg_catalog"."default",
  "difficulty_level" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."courses" OWNER TO "postgres";
COMMENT ON COLUMN "public"."courses"."course_id" IS 'Уникальный идентификатор курса';
COMMENT ON COLUMN "public"."courses"."title" IS 'Название курса';
COMMENT ON COLUMN "public"."courses"."description" IS 'Подробное описание программы курса';
COMMENT ON COLUMN "public"."courses"."difficulty_level" IS 'Уровень сложности: Начальный, Средний, Продвинутый';

-- ----------------------------
-- Records of courses
-- ----------------------------
BEGIN;
INSERT INTO "public"."courses" ("course_id", "title", "description", "difficulty_level") VALUES (0, 'Основы SQL и реляционных баз данных', 'Изучите основы SQL и работу с реляционными базами данных.', 'Начальный');
INSERT INTO "public"."courses" ("course_id", "title", "description", "difficulty_level") VALUES (1, 'Алгоритмы и структуры данных', 'Познакомьтесь с основными алгоритмами и структурами данных.', 'Средний');
INSERT INTO "public"."courses" ("course_id", "title", "description", "difficulty_level") VALUES (2, 'Микросервисная архитектура', 'Освойте принципы проектирования и развертывания микросервисов.', 'Продвинутый');
INSERT INTO "public"."courses" ("course_id", "title", "description", "difficulty_level") VALUES (3, 'Введение в DevOps', 'Научитесь автоматизировать процессы разработки и развертывания.', 'Начальный');
INSERT INTO "public"."courses" ("course_id", "title", "description", "difficulty_level") VALUES (4, 'Безопасность веб-приложений', 'Изучите основы защиты веб-приложений от угроз.', 'Средний');
INSERT INTO "public"."courses" ("course_id", "title", "description", "difficulty_level") VALUES (5, 'Распределённые системы', 'Понимание принципов работы распределённых систем и их масштабирования.', 'Продвинутый');
COMMIT;

-- ----------------------------
-- Table structure for feedback_tickets
-- ----------------------------
DROP TABLE IF EXISTS "public"."feedback_tickets";
CREATE TABLE "public"."feedback_tickets" (
  "ticket_id" int4 NOT NULL,
  "assignment_id" int4,
  "competency_id" int4,
  "description" text COLLATE "pg_catalog"."default",
  "status" varchar(50) COLLATE "pg_catalog"."default",
  "priority" int4
)
;
ALTER TABLE "public"."feedback_tickets" OWNER TO "postgres";
COMMENT ON COLUMN "public"."feedback_tickets"."ticket_id" IS 'Уникальный ID для отслеживания задачи To Do';
COMMENT ON COLUMN "public"."feedback_tickets"."assignment_id" IS 'Задание, в котором была допущена ошибка';
COMMENT ON COLUMN "public"."feedback_tickets"."competency_id" IS 'Навык/тема, которую нужно повторить';
COMMENT ON COLUMN "public"."feedback_tickets"."description" IS 'Четкое описание задачи (напр., ''Исправить JOIN-операцию в файле X'')';
COMMENT ON COLUMN "public"."feedback_tickets"."status" IS 'Состояние задачи: To Do, In Progress, Done';
COMMENT ON COLUMN "public"."feedback_tickets"."priority" IS 'Приоритет задачи (1 - высокий, 3 - низкий)';

-- ----------------------------
-- Records of feedback_tickets
-- ----------------------------
BEGIN;
INSERT INTO "public"."feedback_tickets" ("ticket_id", "assignment_id", "competency_id", "description", "status", "priority") VALUES (0, 0, 0, 'Отличная работа! Запрос корректный и соответствует ТЗ.', 'Done', 1);
INSERT INTO "public"."feedback_tickets" ("ticket_id", "assignment_id", "competency_id", "description", "status", "priority") VALUES (1, 1, 0, 'Хорошо, но можно использовать ILIKE для case-insensitive поиска.', 'Done', 2);
INSERT INTO "public"."feedback_tickets" ("ticket_id", "assignment_id", "competency_id", "description", "status", "priority") VALUES (2, 2, 0, 'В запросе опечатка: WHER вместо WHERE. Исправьте и отправьте повторно.', 'To Do', 1);
COMMIT;

-- ----------------------------
-- Table structure for modules
-- ----------------------------
DROP TABLE IF EXISTS "public"."modules";
CREATE TABLE "public"."modules" (
  "module_id" int4 NOT NULL,
  "course_id" int4,
  "title" varchar(255) COLLATE "pg_catalog"."default",
  "order_index" int4
)
;
ALTER TABLE "public"."modules" OWNER TO "postgres";
COMMENT ON COLUMN "public"."modules"."module_id" IS 'Уникальный идентификатор модуля';
COMMENT ON COLUMN "public"."modules"."course_id" IS 'Идентификатор курса, к которому относится модуль';
COMMENT ON COLUMN "public"."modules"."title" IS 'Название модуля';
COMMENT ON COLUMN "public"."modules"."order_index" IS 'Порядковый номер модуля внутри курса';

-- ----------------------------
-- Records of modules
-- ----------------------------
BEGIN;
INSERT INTO "public"."modules" ("module_id", "course_id", "title", "order_index") VALUES (0, 0, 'Основы SQL', 1);
COMMIT;

-- ----------------------------
-- Table structure for progress
-- ----------------------------
DROP TABLE IF EXISTS "public"."progress";
CREATE TABLE "public"."progress" (
  "progress_id" int4 NOT NULL,
  "student_id" int4,
  "task_id" int4,
  "status" varchar(255) COLLATE "pg_catalog"."default",
  "completion_date" date,
  "score" float4
)
;
ALTER TABLE "public"."progress" OWNER TO "postgres";
COMMENT ON COLUMN "public"."progress"."progress_id" IS 'Уникальный идентификатор записи о прогрессе';
COMMENT ON COLUMN "public"."progress"."student_id" IS 'Ссылка на студента';
COMMENT ON COLUMN "public"."progress"."task_id" IS 'Ссылка на выполненное задание';
COMMENT ON COLUMN "public"."progress"."status" IS 'Статус выполнения: ''Completed'', ''Failed'', ''In Progress''';
COMMENT ON COLUMN "public"."progress"."completion_date" IS 'Дата и время завершения задания';
COMMENT ON COLUMN "public"."progress"."score" IS 'Полученный балл за задание';

-- ----------------------------
-- Records of progress
-- ----------------------------
BEGIN;
INSERT INTO "public"."progress" ("progress_id", "student_id", "task_id", "status", "completion_date", "score") VALUES (0, 0, 0, 'Completed', '2025-11-30', 10);
INSERT INTO "public"."progress" ("progress_id", "student_id", "task_id", "status", "completion_date", "score") VALUES (1, 0, 1, 'Completed', '2025-11-30', 9.5);
INSERT INTO "public"."progress" ("progress_id", "student_id", "task_id", "status", "completion_date", "score") VALUES (2, 0, 2, 'Failed', '2025-11-30', 4);
COMMIT;

-- ----------------------------
-- Table structure for students
-- ----------------------------
DROP TABLE IF EXISTS "public"."students";
CREATE TABLE "public"."students" (
  "student_id" int4 NOT NULL,
  "first_name" varchar(255) COLLATE "pg_catalog"."default",
  "last_name" varchar(255) COLLATE "pg_catalog"."default",
  "email" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "registration_date" date
)
;
ALTER TABLE "public"."students" OWNER TO "postgres";
COMMENT ON COLUMN "public"."students"."student_id" IS 'Уникальный идентификатор студента';
COMMENT ON COLUMN "public"."students"."first_name" IS 'Имя студента';
COMMENT ON COLUMN "public"."students"."last_name" IS 'Фамилия студента';
COMMENT ON COLUMN "public"."students"."email" IS 'Электронная почта студента';
COMMENT ON COLUMN "public"."students"."password" IS 'Пароль';
COMMENT ON COLUMN "public"."students"."registration_date" IS 'Дата регистрации на платформе';

-- ----------------------------
-- Records of students
-- ----------------------------
BEGIN;
INSERT INTO "public"."students" ("student_id", "first_name", "last_name", "email", "password", "registration_date") VALUES (0, 'Иван', 'Иванов', 'ivan@example.com', '$2a$10$c83aASYEZCKupTWCLSkVxO0d.SCFObdtg1s8iXXcK4IaR2fpUav0S', '2025-11-30');
COMMIT;

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS "public"."tasks";
CREATE TABLE "public"."tasks" (
  "task_id" int4 NOT NULL,
  "module_id" int4,
  "title" varchar(255) COLLATE "pg_catalog"."default",
  "task_type" varchar(255) COLLATE "pg_catalog"."default",
  "points_value" int4
)
;
ALTER TABLE "public"."tasks" OWNER TO "postgres";
COMMENT ON COLUMN "public"."tasks"."task_id" IS 'Уникальный ID конкретного задания/урока';
COMMENT ON COLUMN "public"."tasks"."module_id" IS 'Ссылка на модуль, к которому относится задание';
COMMENT ON COLUMN "public"."tasks"."title" IS 'Название задания';
COMMENT ON COLUMN "public"."tasks"."task_type" IS 'Тип контента: Video, Quiz, Reading, Assignment';
COMMENT ON COLUMN "public"."tasks"."points_value" IS 'Количество баллов за успешное выполнение';

-- ----------------------------
-- Records of tasks
-- ----------------------------
BEGIN;
INSERT INTO "public"."tasks" ("task_id", "module_id", "title", "task_type", "points_value") VALUES (0, 0, 'Создание таблицы', 'code', 10);
INSERT INTO "public"."tasks" ("task_id", "module_id", "title", "task_type", "points_value") VALUES (1, 0, 'Выборка данных', 'code', 10);
INSERT INTO "public"."tasks" ("task_id", "module_id", "title", "task_type", "points_value") VALUES (2, 0, 'Ошибочный запрос', 'code', 10);
INSERT INTO "public"."tasks" ("task_id", "module_id", "title", "task_type", "points_value") VALUES (3, 0, 'Типы данных в SQL', 'code', 10);
INSERT INTO "public"."tasks" ("task_id", "module_id", "title", "task_type", "points_value") VALUES (4, 0, 'Первичный и внешний ключ', 'code', 10);
COMMIT;

-- ----------------------------
-- Primary Key structure for table assigments
-- ----------------------------
ALTER TABLE "public"."assigments" ADD CONSTRAINT "assigments_pkey" PRIMARY KEY ("assignment_id");

-- ----------------------------
-- Primary Key structure for table competencies
-- ----------------------------
ALTER TABLE "public"."competencies" ADD CONSTRAINT "competencies_pkey" PRIMARY KEY ("competency_id");

-- ----------------------------
-- Primary Key structure for table courses
-- ----------------------------
ALTER TABLE "public"."courses" ADD CONSTRAINT "courses_pkey" PRIMARY KEY ("course_id");

-- ----------------------------
-- Primary Key structure for table feedback_tickets
-- ----------------------------
ALTER TABLE "public"."feedback_tickets" ADD CONSTRAINT "feedback_tickets_pkey" PRIMARY KEY ("ticket_id");

-- ----------------------------
-- Primary Key structure for table modules
-- ----------------------------
ALTER TABLE "public"."modules" ADD CONSTRAINT "modules_pkey" PRIMARY KEY ("module_id");

-- ----------------------------
-- Primary Key structure for table progress
-- ----------------------------
ALTER TABLE "public"."progress" ADD CONSTRAINT "progress_pkey" PRIMARY KEY ("progress_id");

-- ----------------------------
-- Primary Key structure for table students
-- ----------------------------
ALTER TABLE "public"."students" ADD CONSTRAINT "students_pkey" PRIMARY KEY ("student_id");

-- ----------------------------
-- Primary Key structure for table tasks
-- ----------------------------
ALTER TABLE "public"."tasks" ADD CONSTRAINT "tasks_pkey" PRIMARY KEY ("task_id");

-- ----------------------------
-- Foreign Keys structure for table assigments
-- ----------------------------
ALTER TABLE "public"."assigments" ADD CONSTRAINT "student_id" FOREIGN KEY ("student_id") REFERENCES "public"."students" ("student_id") ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "public"."assigments" ADD CONSTRAINT "task_id" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("task_id") ON DELETE RESTRICT ON UPDATE RESTRICT;

-- ----------------------------
-- Foreign Keys structure for table competencies
-- ----------------------------
ALTER TABLE "public"."competencies" ADD CONSTRAINT "student_id" FOREIGN KEY ("student_id") REFERENCES "public"."students" ("student_id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table feedback_tickets
-- ----------------------------
ALTER TABLE "public"."feedback_tickets" ADD CONSTRAINT "assignment_id" FOREIGN KEY ("assignment_id") REFERENCES "public"."assigments" ("assignment_id") ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "public"."feedback_tickets" ADD CONSTRAINT "competency_id" FOREIGN KEY ("competency_id") REFERENCES "public"."competencies" ("competency_id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table modules
-- ----------------------------
ALTER TABLE "public"."modules" ADD CONSTRAINT "course_id" FOREIGN KEY ("course_id") REFERENCES "public"."courses" ("course_id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table progress
-- ----------------------------
ALTER TABLE "public"."progress" ADD CONSTRAINT "student_id" FOREIGN KEY ("student_id") REFERENCES "public"."students" ("student_id") ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "public"."progress" ADD CONSTRAINT "task_id" FOREIGN KEY ("task_id") REFERENCES "public"."tasks" ("task_id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table tasks
-- ----------------------------
ALTER TABLE "public"."tasks" ADD CONSTRAINT "module_id" FOREIGN KEY ("module_id") REFERENCES "public"."modules" ("module_id") ON DELETE RESTRICT ON UPDATE CASCADE;
