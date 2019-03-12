CREATE DATABASE db_amex01;
CREATE TABLE tbl_todo01 (
	todo_id serial not null,
    task_string varchar( 256),
	posted_at timestamp,
    done boolean
);

INSERT INTO tbl_todo01 (task_string, done, posted_at)
VALUES
 ('hi there00', False, NOW()),
 ('hi there01', False,  NOW()),
 ('hi there02', False,  NOW()),
 ('hi there03', False,  NOW());



CREATE DATABASE db_amex02;
CREATE TABLE tbl_todo01 (
	todo_id serial not null,
    task_string varchar( 256),
	posted_at timestamp,
    done boolean
);

INSERT INTO tbl_todo01 (task_string, done, posted_at)
VALUES
 ('hi there00', False, NOW()),
 ('hi there01', False,  NOW()),
 ('hi there02', False,  NOW()),
 ('hi there03', False,  NOW());


