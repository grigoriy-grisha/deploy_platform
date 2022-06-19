CREATE TABLE project (
  id serial not null unique,
  name varchar(255) not null,
  command varchar(255) not null
);
