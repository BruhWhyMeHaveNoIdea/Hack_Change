-- Down migration: drop schema created by 001_create_schema.up.sql
BEGIN;

DROP TABLE IF EXISTS feedback_tickets CASCADE;
DROP TABLE IF EXISTS competencies CASCADE;
DROP TABLE IF EXISTS assignments CASCADE;
DROP TABLE IF EXISTS progress CASCADE;
DROP TABLE IF EXISTS students CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;
DROP TABLE IF EXISTS modules CASCADE;
DROP TABLE IF EXISTS courses CASCADE;

-- Drop enum types created in the up migration
DO $$ BEGIN
	IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'feedback_status') THEN
		DROP TYPE feedback_status;
	END IF;
	IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'assignment_status') THEN
		DROP TYPE assignment_status;
	END IF;
	IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'progress_status') THEN
		DROP TYPE progress_status;
	END IF;
END $$;

COMMIT;
