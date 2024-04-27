create table users (
    id bigserial primary key,
    oauth_sub varchar unique not null,
    name varchar not null
);

create table posts (
    id bigserial primary key,
    author_id bigint not null references users(id),
    title varchar not null,
    body varchar not null,
    created timestamptz not null default now()
);
