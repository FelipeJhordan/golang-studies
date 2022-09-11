CREATE TABLE publications(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    authorId int noT null,
    FOREIGN KEY (authorId)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0,

    createdAt timestamp default current_timestamp
) ENGINE=INNODB;