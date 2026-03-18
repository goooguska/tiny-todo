CREATE TABLE IF NOT EXISTS tasks(
    id uuid PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);