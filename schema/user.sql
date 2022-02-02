create TYPE user_role AS ENUM ('admin','customer');
CREATE TABLE users(
    id uuid not null primary key,
    name varchar(100) not null,
    password text not null,
    role user_role default 'customer',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

insert into users (id, name, password,role) values
('882c4b8e-6c72-4f32-ac19-2f1431c17c0e','admin','$2a$10$QT03WTl5/6gkhuVPodBD6ebJ.y54PsxJMQWekvjtzRio0Pv95oOty','admin'),
('f0ed9709-c6ef-48b2-8beb-0f88b60d6c83','aji','$2a$10$vuA/29AVCphLnyL4cQYfoeya8dEGuFPjQNHk/XsNL2D8T4CbvCYbe','customer'),
('4d1d51fe-5a6c-4eb7-92a7-38ed97252881','tia','$2a$10$AN2knlxbSmJNOfNPGXAQSuOXCdFIR8rTHFYJvKWoA1UjRLjNgW1tW','customer');

select * from users;
-- drop table users;
-- drop type user_role;


CREATE INDEX name ON users(name);

SELECT name from users where name='aji';
select exists(select name from users where name='aji');
select exists(select name from users where name='aji' and password='secret');


select * from users;
SELECT exists(id) FROM products WHERE id='73e12106-207e-4693-9c0d-3147d6ab606a';

select exists(select id from products where id='73e12106-207e-4693-9c0d-3147d6ab606a');

select name from users where name='aji' and password='aji-secret';
select exists(select name from users where name='aji');

SELECT
    table_name,
    column_name,
    data_type
FROM
    information_schema.columns
WHERE
        table_name = 'users';