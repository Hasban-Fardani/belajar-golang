create table if not exists `users` (
	`id` varchar(7) not null unique,
    `name` varchar(255) not null,
    `email` varchar(255) not null,
    `password` varchar(16) not null,
    `high score` int(5) not null default 0
) engine=InnoDB default charset=latin1;

ALTER TABLE `users` ADD PRIMARY KEY (`id`);
