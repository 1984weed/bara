BEGIN;

-- select 'drop table "' || tablename || '" cascade;' from pg_tables;

INSERT INTO code_languages (id, name, slug) VALUES (1 ,'JavaScript', 'javascript');

INSERT INTO problems (slug, title, description, function_name, output_type, author_id)
    VALUES ('two-sum', 'Two Sum', 'test-description you have to create a code', 'twoSum', 'int[]', 1);

INSERT INTO problem_args (problem_id, order_no, name, var_type) VALUES
    (currval(pg_get_serial_sequence('problems', 'id')), 1, 'nums', 'int[]');
INSERT INTO problem_args (problem_id, order_no, name, var_type) VALUES
    (currval(pg_get_serial_sequence('problems', 'id')), 2, 'target', 'int');

INSERT INTO problem_testcases (problem_id, input_text, output_text) VALUES
    (currval(pg_get_serial_sequence('problems', 'id')), E'[2,7,11,15]\n9', '[0,1]');
INSERT INTO problem_testcases (problem_id, input_text, output_text) VALUES
    (currval(pg_get_serial_sequence('problems', 'id')), E'[3,2,4]\n6', '[1,2]');

INSERT INTO contests(slug, title, start_time) VALUES
    ('test-contest-20200229', 'First contest', '2020-02-29 10:00')

COMMIT;