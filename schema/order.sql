create type status as enum ('success','pending','cancel','deny','expired','settlement');

CREATE TABLE orders(
    id serial primary key,
    code varchar(100) not null,
    product_id uuid not null,
    name varchar(200),
    qty smallint not null,
    total_price double precision not null,
    status status default 'pending',
--     image varchar null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
            REFERENCES products(id)
);
create index on orders (product_id,code);

insert into orders (code, product_id, name, qty, total_price) VALUES
('code-1','d90f8110-039a-47f4-a164-37d807f77ab5','Ipad Pro Max',3,1352.79),
('code-2','73e12106-207e-4693-9c0d-3147d6ab606a','Macbook M1 Pro Max',3,2399.76),
('code-3','73e12106-207e-4693-9c0d-3147d6ab606a','Iphone 12 Pro Max',3,1499.73);

select * from orders;
drop table orders;



CREATE TABLE products(
                         id uuid primary key not null,
                         name varchar(200),
                         stock smallint not null default 0,
                         price double precision not null,
--     image varchar null,
                         created_at timestamp default current_timestamp,
                         updated_at timestamp default current_timestamp
);

insert into products(id, name, stock, price) VALUES
                                                 ('73e12106-207e-4693-9c0d-3147d6ab606a','Iphone 12 Pro Max',51,499.91),
                                                 ('44c22cb3-ff6c-4043-8c79-8a5506ce11e9','Macbook M1 Pro Max',52,799.92),
                                                 ('d90f8110-039a-47f4-a164-37d807f77ab5','Ipad Pro Max',53,450.93);

select * from products;
drop table products;
-- create or replace function when_order_is_expired()
--     returns trigger
--     language plpgsql
-- as
-- $$
-- -- declare
-- --     status_func bool;
-- BEGIN
--     if  new.created_at != current_timestamp then
--         update products set stock = stock from orders where products.id=orders.product_id;
--     end if;
-- end;
-- $$;

