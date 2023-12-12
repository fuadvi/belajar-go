use belajar_gorm;

create table users(
    id varchar(100) not null,
    password varchar(100) not null,
    name varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;

select *
from users;

create table wallets(
    id varchar(100) not null,
    user_id varchar(100) not null,
    balance bigint not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (user_id) references users (id)
)engine  =InnoDB;

select *
from wallets;