drop schema if exists carthago cascade;
create schema carthago;
create extension if not exists pgcrypto;

create table if not exists carthago.product (
    id uuid primary key default gen_random_uuid(),
    description varchar(1000),
    metadata json
);

create table if not exists carthago.batch (
    id uuid primary key default gen_random_uuid(),
    metadata json
);

create table if not exists carthago.order (
    id uuid primary key default gen_random_uuid(),
    metadata json
);

create table if not exists carthago.supplier (
    id uuid primary key default gen_random_uuid(),
    name varchar(100)
);
