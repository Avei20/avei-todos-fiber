-- Altering column todos and delete foreign key
ALTER TABLE todos
DROP COLUMN parent_id,
DROP CONSTRAINT fk_todos_parent_id;
