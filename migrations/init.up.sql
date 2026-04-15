
create schema if not exists "user";

create table if not exists "user"."user" (
    id bigint primary key,
    email text,
    first_name text,
    last_name text,
    avatar_url text,
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

create table if not exists "user".pick_up_points(
    id bigserial primary key,
    user_id bigint not null references "user"."user"(id),
    pick_up_point_id bigint not null
);
