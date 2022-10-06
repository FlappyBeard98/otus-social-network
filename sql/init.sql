create database social_network;

create table social_network.auth
(
    user_id    bigint auto_increment
        primary key,
    login      varchar(50)  not null,
    password   varchar(100) not null,
    last_login datetime     null
);



create table social_network.profile
(
    id         bigint auto_increment
        primary key,
    user_id    bigint                       null,
    first_name varchar(100)                 null,
    last_name  varchar(150)                 null,
    age        int                          null,
    gender     bit                          null,
    city       varchar(50)                  null,
    hobbies    text                         null,
    constraint profile_auth_user_id_fk
        foreign key (user_id) references social_network.auth (user_id)
);

create table social_network.friends
(
    left_user_id  bigint not null,
    right_user_id bigint not null,
    primary key (left_user_id, right_user_id),
    constraint friends_auth_left_user_id_fk
        foreign key (left_user_id) references social_network.auth (user_id),
    constraint friends_auth_right_user_id_fk
        foreign key (right_user_id) references social_network.auth (user_id)
);




