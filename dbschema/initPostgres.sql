create table polls (
                       poll_id varchar(255) primary key not null,
                       chat_id bigint not null,
                       message_id int not null
);

create table users (
                       chat_id bigint not null primary key ,
                       stage varchar(40) not null
);