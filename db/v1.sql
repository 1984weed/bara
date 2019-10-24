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

DROP TABLE IF EXISTS problems cascade;

CREATE TABLE problems (
  id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  title VARCHAR(300) NOT NULL,
  description TEXT NOT NULL,
  function_name VARCHAR(255),
  language_id INTEGER NOT NULL,
  author_id INTEGER,
  output_type args_t NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT problems_slug_unique UNIQUE(slug)
);
ALTER TABLE problems ADD CONSTRAINT fk_problems_code_languages FOREIGN KEY (language_id) REFERENCES code_languages (id);


DROP TABLE IF EXISTS problem_args cascade;

CREATE TABLE problem_args (
  id SERIAL PRIMARY KEY,
  problem_id INTEGER NOT NULL,
  order_no INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  var_type args_t NOT NULL
);

ALTER TABLE problem_args ADD CONSTRAINT fk_problem_args_problems FOREIGN KEY (problem_id) REFERENCES problems (id);

DROP TABLE IF EXISTS problem_testcases cascade;

CREATE TABLE problem_testcases (
  id SERIAL PRIMARY KEY,
  problem_id INTEGER NOT NULL,
  input_text TEXT,
  output_text TEXT NOT NULL
);

ALTER TABLE problem_testcases ADD CONSTRAINT fk_problem_testcases_problems FOREIGN KEY (problem_id) REFERENCES problems (id);

COMMIT;