
create schema if not exists "user";

create table if not exists "user"."user" (
    id bigint primary key,
    email text not null,
    first_name text,
    last_name text,
    phone text,
    avatar_url text,
    created_at timestamptz not null default now()
);

create table if not exists "user".vendor (
    id bigserial primary key,
    user_id bigint not null references "user"."user"(id),
    shop_name text not null,
    description text,
    created_at timestamptz not null default now()
);

create table if not exists "user".address (
    id bigserial primary key,
    user_id bigint not null references "user"."user"(id),
    country text,
    city text,
    street text,
    postal_code text,
    created_at timestamptz not null default now()
);