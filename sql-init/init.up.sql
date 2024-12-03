CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    password_hash varchar(255) not null,
    email varchar(255) not null unique
);

CREATE TABLE course_lists (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    image_url VARCHAR(255),
    course_url VARCHAR(255)
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references course_lists (id) on delete cascade not null
);

CREATE TABLE course_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    image_url VARCHAR(255),
    course_url VARCHAR(255)
);


CREATE TABLE lists_items
(
    id      serial                                           not null unique,
    item_id int references course_items (id) on delete cascade not null,
    list_id int references course_lists (id) on delete cascade not null
);