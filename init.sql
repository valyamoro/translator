CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS dictionaries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS words (
    id SERIAL PRIMARY KEY,
    word VARCHAR(100) NOT NULL,
    translation_word VARCHAR(100) NOT NULL,
    dictionary_id INT REFERENCES dictionaries(id) ON DELETE CASCADE,
    word_language_code INT,
    translation_word_language_code INT
);

INSERT INTO users (username, password) VALUES ('testuser', 'password123');
INSERT INTO dictionaries (name, description, user_id) VALUES ('My Dictionary', 'A test dictionary', 1);
INSERT INTO words (word, translation_word, dictionary_id, word_language_code, translation_word_language_code)
VALUES ('Hello', 'Привет', 1, 1, 2);
