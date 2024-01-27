CREATE TABLE roles
(
    id uuid primary key default uuid_generate_v4(),
    name varchar(255) not null unique
);