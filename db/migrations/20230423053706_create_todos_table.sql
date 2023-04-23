-- migrate:up
create table todos (
  id int not null primary key auto_increment,
  task varchar(255) not null,
  description text
);

-- migrate:down
drop table todos;
