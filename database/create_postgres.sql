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

INSERT INTO cars_for_sale ("id","brand","model","year","img","color","mileage","category","price") VALUES (1,'Hyundai','i30',2016,'/img/cars/hyundai_i30.png','Azul',0,'hatch',81040.00);
INSERT INTO cars_for_sale ("id","brand","model","year","img","color","mileage","category","price") VALUES (2,'Honda','fit',2019,'/img/cars/honda_fit.png','Vermelho',0,'hatch',76035.00);
INSERT INTO cars_for_sale ("id","brand","model","year","img","color","mileage","category","price") VALUES (3,'Toyota','yaris',2019,'/img/cars/toyota_yaris.png','Branco',0,'hatch',84056.00);
INSERT INTO cars_for_sale ("id","brand","model","year","img","color","mileage","category","price") VALUES (4,'Volkswagen','golf',2017,'/img/cars/volkswagen_golf.png','Branco',0,'hatch',79011.00);
