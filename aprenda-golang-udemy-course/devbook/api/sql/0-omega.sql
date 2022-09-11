CREATE DATABASE IF NOT EXISTS devbook;


use devbook;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

CREATE TABLE users(
    id int AUTO_INCREMENT primary key,
    name varchar(50) not null,
    nick varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(100) not null unique,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB


select * from users;

CREATE TABLE followers(
	userId int not null,
	foreign key (userId)
	references users(id)
	on delete cascade,
	
	followerId int not null,
	foreign key (followerId)
	references users(id)
	on delete cascade,
	
	primary key (userId, followerId)
) ENGINE=INNODB;


insert into users (name, nick, email, password) 
values 
("Usuario 1", "usuario", "usuario@uol.com", "$2a$10$7woWhEpM9f5JGQ0000sAP.20b7Xpjg2CTKWgDKx9cz8QQLnFVRBzW"),
("Usuario 2", "usuario2", "usuario2@uol.com", "$2a$10$X112JmXYQCahXu6PCMKtneYeLCLPxbHXKf7qYWwoNR9OoiF3DtAX/q"),
("admin", "admin", "admin@gmail.com.com", "$2a$10$X112JmXYQCahXu6PCMKtneYeLCLPxbHXKf7qYWwoNR9OoiFDtAX/q");

insert into followers(userId, followerId) 
values 
(1, 2),
(2, 1),
(3, 1);