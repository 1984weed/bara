INSERT INTO questions (slug, title, description, function_name, arg_id, language_id, author_id, created_at, updated_at)
    VALUES ('test-slug', 'test title', 'test-description you have to create a code', 'helloWorld', 1, 1, 1, now(), now());

INSERT INTO question_args (question_id, order_no, name, type, created_at, updated_at) VALUES
    (1, 1, 'nums', 'num', now(), now());
INSERT INTO question_args (question_id, order_no, name, type, created_at, updated_at) VALUES
    (1, 2, 'target', 'num', now(), now());

INSERT INTO question_testcases (question_id, input_text, output_text, created_at, updated_at) VALUES
    (1, '[2,7,11,15]\n9', '[0,1]', now(), now());
INSERT INTO question_testcases (question_id, input_text, output_text, created_at, updated_at) VALUES
    (1, '[3,2,4]\n6', '[1,2]', now(), now());