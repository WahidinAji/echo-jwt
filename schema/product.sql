CREATE TABLE products(
    id uuid primary key not null,
    name varchar(200),
    stock smallint not null default 0,
    price numeric(6,2),
--     image varchar null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

insert into products(id, name, stock, price) VALUES
('73e12106-207e-4693-9c0d-3147d6ab606a','Iphone 12 Pro Max',51,499.91),
('44c22cb3-ff6c-4043-8c79-8a5506ce11e9','Macbook M1 Pro Max',52,799.92),
('d90f8110-039a-47f4-a164-37d807f77ab5','Ipad Pro Max',53,450.93);

select * from products;
-- drop table products;