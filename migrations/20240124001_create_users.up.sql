CREATE TABLE users
(
    id uuid primary key default uuid_generate_v4(),
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);