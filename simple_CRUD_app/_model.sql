create database classifieds;

drop table if exists categories cascade;
drop table if exists classifieds cascade;

create table categories (
	category_id serial not null primary key,
	name character varying(64)
);

create table classifieds (
	classified_id serial not null primary key,
	title character varying(256) not null,
	price float not null,
	created_at timestamptz default current_timestamp,
	category_id int not null references categories (category_id)
);

--ALTER TABLE classifieds
--  ADD price float not null;
--MOCK DATA

insert into categories (name) values ('avtomabillar'), ('aloqa vositalari'), ('mebellar');

insert into 
	classifieds (title, price, category_id) 
values ('kompyuter Asus sotiladi', 12000.00, 2)
returning
	*
;

-- QUERIES
--- GET ALL
select
	cl.classified_id,
	cl.title,
	cl.price,
	cl.created_at,
	ct.category_id,
	ct.name
from classifieds cl
join categories as ct using (category_id)
;	

--- GET ONCE
select
	cl.classified_id,
	cl.title,
	cl.price,
	cl.created_at,
	ct.category_id,
	ct.name
from classifieds cl
join categories as ct using (category_id)
where cl.classified_id = 1
;

--- DELETE
delete from 
	classifieds
where
	classified_id = 3
returning 
	*
;

--- UPDATE
update classifieds
set 
    title = coalesce('Avtomabil x5 sotiladi', title),
    price = coalesce(null, price),
    category_id = coalesce(1, category_id)
where classified_id = 1
returning
	*
;
