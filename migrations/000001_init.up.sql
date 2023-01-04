CREATE TABLE users (
    id            serial       not null unique,
    user_id       varchar(50)  not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    first_name    varchar(10),
    last_name     varchar(30),
    date_of_birth date
);
