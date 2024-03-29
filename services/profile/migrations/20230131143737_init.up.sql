create schema if not exists profiles;

create table if not exists profiles.auth
(
    user_id    BIGSERIAL
    primary key,
    login      varchar(250)  not null,
    password  varchar(150) not null,
    constraint login
    unique (login)
);


create table if not exists profiles.friends
(
    user_id  bigint not null,
    friend_user_id bigint not null,
    primary key (user_id, friend_user_id),
    constraint friends_auth_left_user_id_fk
    foreign key (user_id) references profiles.auth (user_id),
    constraint friends_auth_right_user_id_fk
    foreign key (friend_user_id) references profiles.auth (user_id)
);

create table if not exists profiles.profiles
(
    user_id    bigint       not null
    primary key,
    first_name varchar(100) not null default '',
    last_name  varchar(150) not null default '',
    age        int          not null default 0,
    gender     int          not null default 0,
    city       varchar(50)  not null default '',
    hobbies    text         not null,
    constraint profile_auth_user_id_fk
    foreign key (user_id) references profiles.auth (user_id)
);
