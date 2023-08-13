DROP DATABASE IF EXISTS socialmedia;

CREATE DATABASE socialmedia;

\c socialmedia;

CREATE TABLE users (user_id SERIAl PRIMARY KEY,  username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL);

INSERT INTO users (username, password) VALUES
    ('user1', 'password1'),
    ('user2', 'password2');

CREATE TABLE messages (message_id SERIAL PRIMARY KEY, sender_id INT REFERENCES users(user_id), message_text  VARCHAR(100) NOT NULL);

INSERT INTO messages (sender_id, message_text) VALUES (1, 'hello from 1'),(2, 'hello from 2');