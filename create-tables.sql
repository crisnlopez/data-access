drop table if exists album;
create table album (
  id int auto_increment not null,
  title varchar(128) not null,
  artist varchar(255) not null,
  price decimal(5,2) not null,
  primary key (`id`)
);

insert into album
(title, artist, price)
values
('blue train','john coltrane', 56.99),
('giant steps','john coltrane', 63.99),
('jeru','gerry mulligan', 17.99),
('sarah vaughan','sarah vaughnan', 34.98);
