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

CREATE TABLE public.cars_for_service (
	id SERIAL,
	customer_id int4 NOT NULL,
	car_details varchar(160) NOT NULL,
	mechanic varchar(160) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.customer (
	id SERIAL,
	first_name varchar(50) NOT NULL,
	last_name varchar(50) NOT NULL,
	birth_date date NOT NULL,
	email varchar(100) NULL,
	phone varchar(20) NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.salesperson (
	id SERIAL,
	first_name varchar(50) NOT NULL,
	last_name varchar(50) NULL,
	commission numeric(4, 2) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE public.invoice (
	id SERIAL,
	customer_id int4 NOT NULL,
	salesperson_id int4 NOT NULL,
	car_id int4 NOT NULL,
	amount int4 NOT NULL,
	PRIMARY KEY (id)
);

-- data for table cars_for_sale
INSERT INTO cars_for_sale ("id","brand","model","year","img","color","mileage","category","price") VALUES (1,'Hyundai','i30',2016,'/img/cars/hyundai_i30.png','Azul',0,'hatch',81040.00);
INSERT INTO cars_for_sale ("id","brand","model","year","img","color","mileage","category","price") VALUES (2,'Honda','fit',2019,'/img/cars/honda_fit.png','Vermelho',0,'hatch',76035.00);
INSERT INTO cars_for_sale ("id","brand","model","year","img","color","mileage","category","price") VALUES (3,'Toyota','yaris',2019,'/img/cars/toyota_yaris.png','Branco',0,'hatch',84056.00);
INSERT INTO cars_for_sale ("id","brand","model","year","img","color","mileage","category","price") VALUES (4,'Volkswagen','golf',2017,'/img/cars/volkswagen_golf.png','Branco',0,'hatch',79011.00);

-- data for table customer
INSERT INTO customer ("id","first_name","last_name","birth_date","email","phone") VALUES (1,'Herbert','Olga','1982-11-10','herbertolga@gmail.com.ca','99727701');
INSERT INTO customer ("id","first_name","last_name","birth_date","email","phone") VALUES (2,'Isabela','Cristina','1997-05-05','isabela@gmail.com','99700522');
INSERT INTO customer ("id","first_name","last_name","birth_date","email","phone") VALUES (3,'Caio','Batista','1990-10-06','caio_batista@hotmail.com','99700113');

-- data for table salesperson
INSERT INTO salesperson ("id","first_name","last_name","commission") VALUES (1,'Candido','Martins',4.30);
INSERT INTO salesperson ("id","first_name","last_name","commission") VALUES (2,'Lidia','Alcantara',3.90);
INSERT INTO salesperson ("id","first_name","last_name","commission") VALUES (3,'Maria','Menezes',4.10);
INSERT INTO salesperson ("id","first_name","last_name","commission") VALUES (4,'Rodolfo','Zimmerman',4.20);
