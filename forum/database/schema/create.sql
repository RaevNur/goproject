CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    nickname TEXT,
    email TEXT,
    created_at TEXT,
    password TEXT,
    first_name TEXT,
    last_name TEXT
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY,
    uuid TEXT,
    user_id INTEGER,
    created_at TEXT
);

CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY,
    name TEXT
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY,
    content TEXT,
    user_id INTEGER,
    created_at TEXT
);

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY,
    content TEXT,
    user_id INTEGER,
    created_at TEXT,
    post_id INTEGER
);

CREATE TABLE IF NOT EXISTS questions (
    id INTEGER PRIMARY KEY,
    post_id INTEGER,
    title TEXT,
    views INTEGER
);

CREATE TABLE IF NOT EXISTS likes (
    id INTEGER PRIMARY KEY,
    user_id TEXT,
    post_id TEXT,
    liked INTEGER
);

CREATE TABLE IF NOT EXISTS tags_questions (
    id INTEGER PRIMARY KEY,
    tag_ig INTEGER,
    question_id INTEGER
);