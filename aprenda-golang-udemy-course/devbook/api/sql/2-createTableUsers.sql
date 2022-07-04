use devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int AUTO_INCREMENT primary key,
    name varchar(50) not null,
    nick varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(50) not null unique,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB