CREATE TABLE users
(
    id            serial PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_list
(
    id serial PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE todo_items
(
    id serial PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

CREATE TABLE users_lists
(
    id serial PRIMARY KEY,
    user_id INT NOT NULL,
    list_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE ON UPDATE RESTRICT,
    FOREIGN KEY (list_id) REFERENCES todo_list ON DELETE CASCADE ON UPDATE RESTRICT
);

CREATE TABLE lists_items
(
    id serial PRIMARY KEY,
    list_id INT NOT NULL,
    item_id INT NOT NULL,
    FOREIGN KEY (list_id) REFERENCES todo_list ON DELETE CASCADE ON UPDATE RESTRICT,
    FOREIGN KEY (item_id) REFERENCES todo_items ON DELETE CASCADE ON UPDATE RESTRICT
);