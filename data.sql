create table users
(
    id       bigserial
        primary key,
    username text
        unique,
    password text,
    email    text
        unique,
    avatar   text default 'default.jpg'::text,
    roles    integer[],
    unique (username, email)
);

alter table users
    owner to postgres;

INSERT INTO public.users (id, username, password, email, avatar, roles) VALUES (2, 'admin', '123456', 'admin@gmail.com', 'default.jpg', '{3}');
INSERT INTO public.users (id, username, password, email, avatar, roles) VALUES (5, 'user2', '123456', 'user2@gmail.com', 'default.jpg', '{1}');
INSERT INTO public.users (id, username, password, email, avatar, roles) VALUES (6, 'user1', '123456', 'user1@gmail.com', 'default.jpg', '{1}');
INSERT INTO public.users (id, username, password, email, avatar, roles) VALUES (4, 'teacher2', '123456', 'teacher2@gmail.com', 'default.jpg', '{2}');
INSERT INTO public.users (id, username, password, email, avatar, roles) VALUES (3, 'teacher1', '123456', 'teacher1@gmail.com', 'default.jpg', '{2}');
INSERT INTO public.users (id, username, password, email, avatar, roles) VALUES (1, 'admin2', '123456', 'admin2@gmail.com', 'default.jpg', '{3}');

create table todos
(
    id            bigserial
        primary key,
    name          text   not null,
    description   text,
    user_id       bigint not null,
    todo_group_id integer
        constraint fk_customer
            references group_todos
);

create table group_todos
(
    id      bigserial
        primary key,
    name    text   not null,
    user_id bigint not null
);

alter table group_todos
    owner to postgres;

INSERT INTO public.group_todos (id, name, user_id) VALUES (1, 'Việc Nhà', 1);
INSERT INTO public.group_todos (id, name, user_id) VALUES (2, 'Việc Học', 1);


alter table todos
    owner to postgres;

INSERT INTO public.todos (id, name, description, user_id, todo_group_id) VALUES (1, 'Quét nhà', '4h-5h', 1, 1);
INSERT INTO public.todos (id, name, description, user_id, todo_group_id) VALUES (2, 'Quét sân', '6h-7h', 1, 1);
INSERT INTO public.todos (id, name, description, user_id, todo_group_id) VALUES (3, 'Học Toàn', '5h-6h', 1, 1);


create table rbac_rules
(
    id        serial
        primary key,
    path      varchar(255) not null,
    method    varchar(10)  not null,
    auth_type varchar(20)  not null,
    roles     integer[],
    name      varchar(50),
    status    boolean,
    service   varchar(50)
);

alter table rbac_rules
    owner to postgres;

INSERT INTO public.rbac_rules (id, path, method, auth_type, roles, name, status, service) VALUES (3, '/todo-groups', 'GET', 'ALLOW', '{1}', 'Xem tất cả công việc', true, 'MAIN');
INSERT INTO public.rbac_rules (id, path, method, auth_type, roles, name, status, service) VALUES (1, '/account', 'GET', 'ALLOW_ALL', '', 'Xem 1 thông user', true, 'MAIN');
INSERT INTO public.rbac_rules (id, path, method, auth_type, roles, name, status, service) VALUES (2, '/users/{id}', 'GET', 'ALLOW_ALL', null, 'Xem 1 user', true, 'MAIN');


create table roles
(
    id   serial
        primary key,
    name varchar(50) not null
        unique
);

alter table roles
    owner to postgres;

INSERT INTO public.roles (id, name) VALUES (3, 'ADMIN');
INSERT INTO public.roles (id, name) VALUES (1, 'STUDENT');
INSERT INTO public.roles (id, name) VALUES (2, 'TEACHER');

create table tags
(
    id    bigserial
        primary key,
    name  text not null,
    color text not null
);

alter table tags
    owner to postgres;

create table todo_tags
(
    id      bigserial
        primary key,
    todo_id bigint not null,
    tag_id  bigint not null
);

alter table todo_tags
    owner to postgres;

