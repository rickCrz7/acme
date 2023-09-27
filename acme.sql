drop table if exists invoice_item;
drop table if exists invoice;
drop table if exists product;
drop table if exists customer;

create table customer (
    id serial primary key,
    name varchar(50) not null
);

create table product (
    id serial primary key,
    name text not null,
    price money not null
);

create table invoice (
    id serial primary key,
    customer_id serial not null references customer(id) on delete cascade,
    purchaseDate date not null
);

create table invoice_item (
    id serial primary key,
    invoice_id serial not null references invoice(id) on delete cascade,
    product_id serial not null references product(id) on delete cascade,
    quantity integer not null
);

INSERT INTO customer (name) VALUES ('Acme Corp');
INSERT INTO customer (name) VALUES ('SWAU EDU');

INSERT INTO product (name, price) VALUES ('Computers', 499.99);
INSERT INTO product (name, price) VALUES ('Monitors', 149.99);
INSERT INTO product (name, price) VALUES ('Keyboards', 49.99);

INSERT INTO invoice (customer_id, purchaseDate) VALUES (1, '2023-01-01');
INSERT INTO invoice (customer_id, purchaseDate) VALUES (2, '2023-03-02');

INSERT INTO invoice_item (invoice_id, product_id, quantity) VALUES (1, 1, 100);
INSERT INTO invoice_item (invoice_id, product_id, quantity) VALUES (1, 2, 150);
INSERT INTO invoice_item (invoice_id, product_id, quantity) VALUES (1, 3, 100);

INSERT INTO invoice_item (invoice_id, product_id, quantity) VALUES (2, 1, 50);
INSERT INTO invoice_item (invoice_id, product_id, quantity) VALUES (2, 2, 100);



