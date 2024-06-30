-- add column code and unique into projects table
ALTER TABLE projects
ADD COLUMN code VARCHAR(36) UNIQUE;
