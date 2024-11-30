CREATE TABLE users (
                       id      serial       not null unique,
                       username varchar(255) not null,
                       password varchar(255) not null,
                        email varchar(255) not null unique
);

CREATE TABLE courses (
                         id serial not null unique,
                         name varchar(255) not null,
                         description varchar(255),
                         user_id int references users(id) on delete set null
);
