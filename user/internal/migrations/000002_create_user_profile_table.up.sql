create table if not exists user_profiles (
     id bigint unsigned not null auto_increment,
     user_id bigint unsigned not null,
     first_name varchar(100),
     last_name varchar(100),
     primary key (id),
     foreign key (user_id) references users(id)
);
