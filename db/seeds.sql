INSERT INTO questions (slug, title, description, function_name, arg_id, language_id, author_id, created_at, updated_at)
    VALUES ('test-slug', 'test title', 'test-description you have to create a code', 'helloWorld', 1, 1, 1, now(), now());

INSERT INTO question_args (question_id, order_no, name, type) VALUES
    (2, 1, 'nums', 'num');
INSERT INTO question_args (question_id, order_no, name, type) VALUES
    (2, 2, 'target', 'num');