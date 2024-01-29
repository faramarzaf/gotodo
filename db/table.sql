create table todo
(

    id          bigint primary key AUTO_INCREMENT,
    title       varchar(256) not null,
    description varchar(256) not null,
    done        boolean      not null,
    created_at  timestamp default current_timestamp
)
