create table t_sqlx_gen (
    id int(11) auto_increment primary key,
    name varchar(100) not null default "",
    age tinyint not null default 0,
    gender tinyint not null default 0
) charset=utf8mb4;