INSERT INTO code_languages (id, name, slug) VALUES (1 ,'JavaScript', 'javascript');

INSERT INTO question_args (question_id, order_no, name, type) VALUES
    (1, 1, 'nums', 'num');
INSERT INTO question_args (question_id, order_no, name, type) VALUES
    (1, 2, 'target', 'num');

INSERT INTO question_testcases (question_id, input_text, output_text) VALUES
    (1, E'[2,7,11,15]\n9', '[0,1]');
INSERT INTO question_testcases (question_id, input_text, output_text) VALUES
    (1, E'[3,2,4]\n6', '[1,2]');

INSERT INTO questions (slug, title, description, function_name, output_type, language_id, author_id, created_at, updated_at)
    VALUES ('two-sum', 'Two Sum', 'test-description you have to create a code', 'twoSum', 'int[]', 1, 1, now(), now());