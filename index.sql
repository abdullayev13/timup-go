create table if not exists sms_codes
(
    id           bigserial primary key,
    phone_number text,
    code         text,
    sent_at      timestamp with time zone,
    verified     boolean
);

alter table sms_codes
    owner to postgres;


create table if not exists users
(
    id           bigserial
        primary key,
    fist_name    text,
    last_name    text,
    password     text,
    user_name    text not null
        constraint idx_users_user_name
            unique,
    phone_number text not null
        constraint idx_users_phone_number
            unique,
    address      text,
    photo_url    text,
    birth_date   timestamp with time zone
);

alter table users
    owner to postgres;

create unique index idx_username
    on users (user_name);

create unique index idx_phone_number
    on users (phone_number);


create table if not exists work_categories
(
    id        bigserial
        primary key,
    parent_id bigint references work_categories (id),
    name      text not null
);

alter table work_categories
    owner to postgres;

create unique index idx_parent_id__name
    on work_categories (parent_id, name);


create table if not exists business_profiles
(
    id               bigserial
        primary key,
    user_id          bigint references users (id),
    work_category_id bigint references work_categories (id),
    office_address   text,
    office_name      text,
    experience       bigint,
    bio              text,
    day_offs         text
);

alter table business_profiles
    owner to postgres;


create table if not exists bookings
(
    id          bigserial
        primary key,
    client_id   bigint references users (id),
    date        timestamp with time zone,
    business_id bigint references business_profiles (id)
);

alter table bookings
    owner to postgres;


create table if not exists followings
(
    id          bigserial
        primary key,
    business_id bigint not null references business_profiles (id),
    follower_id bigint not null references users (id),
    created_at  timestamp with time zone
);

alter table followings
    owner to postgres;

create unique index idx_following
    on followings (business_id, follower_id);

