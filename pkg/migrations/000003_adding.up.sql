-- adding column to todos with a foreign key to the same table
ALTER TABLE todos
ADD COLUMN parent_id VARCHAR(36),
ADD CONSTRAINT fk_todos_parent_id FOREIGN KEY (parent_id) REFERENCES todos(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;
