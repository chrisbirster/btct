-- 001_add_owner_id_to_tasks.sql

ALTER TABLE tasks ADD COLUMN owner_id TEXT NOT NULL DEFAULT 'default_user';

-- Assign a specific owner to existing tasks
UPDATE tasks SET owner_id = 'user1' WHERE owner_id = '';

INSERT INTO migrations (count, description)
VALUES (1, 'Add owner_id column to tasks table and set default owner');

