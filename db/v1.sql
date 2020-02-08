BEGIN;
DROP TABLE IF EXISTS users cascade;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  user_name VARCHAR(255) NOT NULL,
  real_name VARCHAR(255),
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  bio VARCHAR(1024),
  image VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT user_email_unique UNIQUE (email),
  CONSTRAINT user_username_unique UNIQUE (user_name)
);

DROP TABLE IF EXISTS contests cascade;

CREATE TABLE contests (
  id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  start_time  TIMESTAMP NOT NULL,
  CONSTRAINT contest_slug_unique UNIQUE (slug)
);

DROP TABLE IF EXISTS contest_problems cascade;

CREATE TABLE contest_problems (
  id SERIAL PRIMARY KEY,
  contest_id INTEGER NOT NULL,
  problem_id INTEGER NOT NULL,
  order_id INTERGER 
);

ALTER TABLE contest_problems ADD CONSTRAINT fk_contest_problems_problem_id FOREIGN KEY (problem_id) REFERENCES problems (id);
ALTER TABLE contest_problems ADD CONSTRAINT fk_contest_problems_contest_id FOREIGN KEY (contest_id) REFERENCES contests (id);

DROP TABLE IF EXISTS code_languages cascade;

CREATE TABLE code_languages (
  id INTEGER PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  slug VARCHAR(255) NOT NULL
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
  author_id INTEGER,
  output_type args_t NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT problems_slug_unique UNIQUE(slug)
);

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

DROP TABLE IF EXISTS problem_user_results cascade;

CREATE TABLE problem_user_results (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  problem_id INTEGER NOT NULL,
  submitted_code TEXT NOT NULL,
  status VARCHAR(10) NOT NULL,
  code_lang_id INTEGER NOT NULL,
  exec_time INTEGER,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE problem_user_results ADD CONSTRAINT fk_problem_user_results_uses FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE problem_user_results ADD CONSTRAINT fk_problem_user_results_problems FOREIGN KEY (problem_id) REFERENCES problems (id);
ALTER TABLE problem_user_results ADD CONSTRAINT fk_problem_user_results_code_languages FOREIGN KEY (code_lang_id) REFERENCES code_languages (id);

COMMIT;