BEGIN;

INSERT INTO users (name, unique_name, password) VALUES ('admin', 'admin', 'admin');
INSERT INTO code_languages (id, name, unique_key) VALUES (1 ,'JavaScript', 'javascript');

COMMIT;