DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  bio VARCHAR(1024),
  image VARCHAR(255),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT user_email_unique UNIQUE (email),
  CONSTRAINT user_username_unique UNIQUE (username)
);

DROP TABLE IF EXISTS questions;

CREATE TABLE questions (
  id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  title VARCHAR(300) NOT NULL,
  description TEXT NOT NULL,
  function_name VARCHAR(255),
  arg_id INTEGER,
  language_id INTEGER,
  author_id INTEGER,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT questions_slug_unique UNIQUE(slug)
);

DROP TYPE IF EXISTS args_t;
CREATE TYPE args_t AS enum('string', 'list', 'num', 'likedList', 'node');

DROP TABLE IF EXISTS question_args;

CREATE TABLE question_args (
  id SERIAL PRIMARY KEY,
  question_id INTEGER NOT NULL,
  order_no INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  type args_t NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  constraint fk_question_args_question FOREIGN KEY (question_id) REFERENCES questions (id)
);

DROP TABLE IF EXISTS question_testcases;

CREATE TABLE question_testcases (
  id SERIAL PRIMARY KEY,
  question_id INTEGER NOT NULL,
  input_text TEXT NOT NULL,
  output_text TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  constraint fk_question_testcases_question FOREIGN KEY (question_id) REFERENCES questions (id)
);

