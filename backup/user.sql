create table users
(
    id bigserial primary key,
    email text not null,
    name text not null,
    password varchar(60) not null,
    role text default 'user'
);