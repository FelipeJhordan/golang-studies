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