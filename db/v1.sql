BEGIN;
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

DROP TABLE IF EXISTS code_languages cascade;
CREATE TABLE code_languages (
  id INTEGER PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  slug  VARCHAR(255) NOT NULL
);



DROP TYPE IF EXISTS args_t cascade;
CREATE TYPE args_t AS enum(
  'string[][]',
  'string[]', 
  'string',
  'int[][]',
  'int[]',
  'int',
  'double[][]',
  'double[]',
  'double');

DROP TABLE IF EXISTS questions cascade;

CREATE TABLE questions (
  id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  title VARCHAR(300) NOT NULL,
  description TEXT NOT NULL,
  function_name VARCHAR(255),
  language_id INTEGER NOT NULL,
  author_id INTEGER,
  output_type args_t NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT questions_slug_unique UNIQUE(slug)
);
ALTER TABLE questions ADD CONSTRAINT fk_questions_code_languages FOREIGN KEY (language_id) REFERENCES code_languages (id);


DROP TABLE IF EXISTS question_args cascade;

CREATE TABLE question_args (
  id SERIAL PRIMARY KEY,
  question_id INTEGER NOT NULL,
  order_no INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  var_type args_t NOT NULL
);

ALTER TABLE question_args ADD CONSTRAINT fk_question_args_questions FOREIGN KEY (question_id) REFERENCES questions (id);

DROP TABLE IF EXISTS question_testcases cascade;

CREATE TABLE question_testcases (
  id SERIAL PRIMARY KEY,
  question_id INTEGER NOT NULL,
  input_text TEXT,
  output_text TEXT NOT NULL
);

ALTER TABLE question_testcases ADD CONSTRAINT fk_question_testcases_questions FOREIGN KEY (question_id) REFERENCES questions (id);

COMMIT;