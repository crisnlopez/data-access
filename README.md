# data-access
Practice connect to a Database with go

## Database structure
create-table.sql has the sql code to create the database will be use
~~~
drop table if exists album;
create table album (
  id int auto_increment not null,
  title varchar(128) not null,
  artist varchar(255) not null,
  price decimal(5,2) not null,
  primary key('id')
);

insert into album (title, artist, price)
values
('blue train','john coltrane',56.99),
('giant steps','john coltrane', 63.99),
('jeru','gerry mulligan', 17.99),
('sarah vaughan','sarah vaughnan', 34.98);
~~~

## DBUSER and DBPASS env
Set your env in the command line.
You will need to pass this envs in mysql.Config{ ... }
~~~
export DBUSER=your_user
export DBPASS=your_password
~~~

## More info
For more info visit [Go tutorial, accessing a relational database](https://go.dev/doc/tutorial/database-access)
