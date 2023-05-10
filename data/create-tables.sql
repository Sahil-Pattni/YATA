-- Drop table if exists
DROP TABLE IF EXISTS todo;

CREATE TABLE todo (
    id INTEGER PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT 0
);