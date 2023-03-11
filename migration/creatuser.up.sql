create table if not exists users(
    id UUID not null primary key,
    username varchar(32),
    password varchar(32),
    phone text not null,
);