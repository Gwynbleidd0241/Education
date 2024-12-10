CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    password_hash varchar(255) not null,
    email varchar(255) not null unique
);

CREATE TABLE course (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    image_url TEXT,
    course_url TEXT,
    profession TEXT
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references course (id) on delete cascade not null
);

CREATE TABLE user_courses (
                              id SERIAL PRIMARY KEY,
                              user_id INT NOT NULL REFERENCES users(id),
                              course_id INT NOT NULL REFERENCES course(id),
                              UNIQUE (user_id, course_id)
);