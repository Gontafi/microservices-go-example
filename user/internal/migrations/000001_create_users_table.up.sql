create table if not exists users (
    id bigint unsigned not null auto_increment,
    username varchar(100) not null,
    email varchar(100),
    passwordHash varchar(255) not null,
    created_at timestamp not null default current_timestamp(),
    updated_at timestamp not null default current_timestamp() on update current_timestamp(),
    primary key (id)
);
