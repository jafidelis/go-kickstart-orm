CREATE TABLE users (
	id serial,
    first_name varchar(60) NOT NULL,
    last_name varchar(60) NOT NULL,
    email varchar(40),
    login varchar(20) NOT NULL,
    password text NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    CONSTRAINT users_pk PRIMARY KEY (id)
);