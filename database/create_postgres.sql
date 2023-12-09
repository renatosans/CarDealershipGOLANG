CREATE DATABASE car_dealership;
\c car_dealership

CREATE TABLE public.cars_for_sale (
	id SERIAL,
	brand varchar(50) NOT NULL,
	model varchar(100) NOT NULL,
	"year" int4 NOT NULL,
	img varchar(255) NULL,
	color varchar(50) NULL,
	mileage int4 NULL,
	category varchar(50) NULL,
	price numeric(10, 2) NOT NULL,
	PRIMARY KEY (id)
);
