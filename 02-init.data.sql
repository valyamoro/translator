INSERT INTO users (username, password) VALUES ('testuser', 'password123');
INSERT INTO dictionaries (name, description, user_id) VALUES ('My Dictionary', 'A test dictionary', 1);
INSERT INTO words (word, translation_word, dictionary_id, word_language_code, translation_word_language_code)
VALUES ('Hello', 'Привет', 1, 1, 2);
