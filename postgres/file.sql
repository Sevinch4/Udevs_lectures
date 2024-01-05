create table users(
    id serial primary key,
    first_name varchar(30),
    last_name varchar(30),
    phone varchar(13)
);

create table orders(
    id serial primary key,
    user_id int references users(id) not null,
    total_sum int
);

create table products(
    id int,
    name varchar(30),
    quantity int,
    price int,
    order_id int references orders(id) not null
);