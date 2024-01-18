CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
)

CREATE TABLE notes (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT,
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
)