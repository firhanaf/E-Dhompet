create database db_edhompet;

use db_edhompet;

create table users(
id varchar(100) primary key unique not null,
username varchar (100) unique not null,
name varchar(100) not null,
address text,
phone varchar(15) unique not null,
email varchar(100) unique not null,
password varchar(100) not null
);

select * from users;

alter table users
MODIFY COLUMN balance decimal default 0;

insert into users(id, username, name, address, phone, email, password)
values("UID-00001", "adi12", "adi", "aceh", "081234", "adi@dai.com", "qwerty"),
("UID-00002", "baba10", "baba", "bandung", "081345", "baba@mail.com", "qwerty123"),
("UID-00003", "charlie23", "charlie", "cikampek", "081456", "char@lie.com", "123qwerty"),
("UID-00004", "deri07", "deri", "depok", "081567", "deri@dare-i.com", "123qwe"),
("UID-00005", "epi123", "epi", "solo", "081678", "epi@ipe.com", "qwe123");

select * from users;

create table transfers(
id int primary key default 238001 unique,
user_id varchar(100) not null,
receiver_userid varchar(100),
amount int not null,
status varchar(10) not null,
transaction_time datetime default current_timestamp,
constraint fk_transfers_users foreign key(user_id) references users(id)
);

insert into transfers(user_id, receiver_userid, amount, status)
values("UID-00001", "UID-00002", 20000, "SUCCESS");

-- saldo pengirim berkurang
update users set
balance = 80000 where id = "UID-00001";

-- saldo penerima bertambah
update users set
balance = 20000 where id = "UID-00002";

select * from transfers;

create table topup(
id int primary key default 823001 unique,
user_id varchar(100),
amount int,
status varchar(100),
transaction_time datetime default current_timestamp,
constraint fk_topup_users foreign key(user_id) references users(id)
);

insert into topup(user_id, amount, status)
values("UID-00001", 100000, "SUCCESS");

-- update isi saldo
update users SET
balance = 100000 where id = "UID-00001";

select * from topup;

SELECT * FROM TRANSFERS;

INSERT INTO TRANSFERS (user_id,receiver_userid, amount, status)
values ("UID-00002", "UID-0003", 50000, "SUCCESS");

insert into transfers(user_id, receiver_userid, amount, status)
values("UID-00001", "UID-00002", 20000, "SUCCESS");


