drop database if exists map;
create database map;

create table country(
    id serial primary key,
    name varchar(50) not null unique
);

create table city(
    id serial primary key,
    country_id int references country(id),
    is_capital boolean default false,
    found_at timestamp not null,
    name varchar(50) not null,
    population int
);
