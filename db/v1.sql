BEGIN;
  DROP TABLE IF EXISTS accounts
  cascade;

CREATE TABLE accounts
(
  id SERIAL,
  compound_id VARCHAR(255) NOT NULL,
  user_id INTEGER NOT NULL,
  provider_type VARCHAR(255) NOT NULL,
  provider_id VARCHAR(255) NOT NULL,
  provider_account_id VARCHAR(255) NOT NULL,
  refresh_token TEXT,
  access_token TEXT,
  access_token_expires TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS sessions
cascade;

CREATE TABLE sessions
(
  id SERIAL,
  user_id INTEGER NOT NULL,
  expires TIMESTAMPTZ NOT NULL,
  session_token VARCHAR(255) NOT NULL,
  access_token VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS users
cascade;

CREATE TABLE users
(
  id SERIAL,
  name VARCHAR(255),
  unique_name VARCHAR(128),
  role VARCHAR(32),
  password VARCHAR(128),
  email VARCHAR(255),
  email_verified TIMESTAMPTZ,
  image VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS verification_requests
cascade;

CREATE TABLE verification_requests
(
  id SERIAL,
  identifier VARCHAR(255) NOT NULL,
  token VARCHAR(255) NOT NULL,
  expires TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX compound_id
  ON accounts(compound_id);

CREATE INDEX provider_account_id
  ON accounts(provider_account_id);

CREATE INDEX provider_id
  ON accounts(provider_id);

CREATE INDEX user_id
  ON accounts(user_id);

CREATE UNIQUE INDEX session_token
  ON sessions(session_token);

CREATE UNIQUE INDEX access_token
  ON sessions(access_token);

CREATE UNIQUE INDEX email
  ON users(email);

CREATE UNIQUE INDEX unique_name
  ON users(unique_name)

CREATE UNIQUE INDEX token
  ON verification_requests(token);


DROP TYPE IF EXISTS args_t
cascade;
CREATE TYPE args_t AS enum
(
  'string[][]',
  'string[]', 
  'string',
  'int[][]',
  'int[]',
  'int',
  'double[][]',
  'double[]',
  'double');

DROP TABLE IF EXISTS problems
cascade;

CREATE TABLE problems
(
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

DROP TABLE IF EXISTS contests
cascade;

CREATE TABLE contests
(
  id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  start_time TIMESTAMP NOT NULL,
  CONSTRAINT contest_slug_unique UNIQUE (slug)
);

DROP TABLE IF EXISTS contest_problems
cascade;

CREATE TABLE contest_problems
(
  id SERIAL PRIMARY KEY,
  contest_id INTEGER NOT NULL,
  problem_id INTEGER NOT NULL,
  order_id INTEGER
);

ALTER TABLE contest_problems ADD CONSTRAINT fk_contest_problems_problem_id FOREIGN KEY (problem_id) REFERENCES problems (id);
ALTER TABLE contest_problems ADD CONSTRAINT fk_contest_problems_contest_id FOREIGN KEY (contest_id) REFERENCES contests (id);

DROP TABLE IF EXISTS contest_problem_user_results
cascade;

CREATE TABLE contest_problem_user_results
(
  id SERIAL PRIMARY KEY,
  contest_id INTEGER NOT NULL,
  problem_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  status VARCHAR(10) NOT NULL,
  exec_time INTEGER,
  created_at TIMESTAMP NOT NULL
);

ALTER TABLE contest_problem_user_results ADD CONSTRAINT fk_contest_problems_problem_id FOREIGN KEY (problem_id) REFERENCES problems (id);
ALTER TABLE contest_problem_user_results ADD CONSTRAINT fk_contest_problems_contest_id FOREIGN KEY (contest_id) REFERENCES contests (id);

DROP TABLE IF EXISTS contest_user_results
cascade;

CREATE TABLE contest_user_results
(
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  contest_id INTEGER NOT NULL,
  ranking INTEGER NOT NULL
);

DROP TABLE IF EXISTS code_languages
cascade;

CREATE TABLE code_languages
(
  id INTEGER PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  unique_key VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX unique_key
  ON code_languages(unique_key);

DROP TABLE IF EXISTS problem_args
cascade;

CREATE TABLE problem_args
(
  id SERIAL PRIMARY KEY,
  problem_id INTEGER NOT NULL,
  order_no INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  var_type args_t NOT NULL
);

ALTER TABLE problem_args ADD CONSTRAINT fk_problem_args_problems FOREIGN KEY (problem_id) REFERENCES problems (id);

DROP TABLE IF EXISTS problem_testcases
cascade;

CREATE TABLE problem_testcases
(
  id SERIAL PRIMARY KEY,
  problem_id INTEGER NOT NULL,
  input_text TEXT,
  output_text TEXT NOT NULL
);

ALTER TABLE problem_testcases ADD CONSTRAINT fk_problem_testcases_problems FOREIGN KEY (problem_id) REFERENCES problems (id);

DROP TABLE IF EXISTS problem_user_results
cascade;

CREATE TABLE problem_user_results
(
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