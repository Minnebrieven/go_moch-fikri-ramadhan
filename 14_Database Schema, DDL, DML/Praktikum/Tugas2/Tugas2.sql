--Tugas 2 

--1.Create database alta_online_shop.
CREATE DATABASE alta_online_shop;


--2.Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
--2.A.Create table user
CREATE TABLE user (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	name VARCHAR(64) NOT NULL,
	address TEXT,
	birthdate DATE,
	status SMALLINT,
	gender CHAR(1),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--2.B.Create table product, product type, operators, product description, payment_method.
--table product
CREATE TABLE product (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	name VARCHAR(255) NOT NULL,
	price INT,
	product_type_id INT,
	product_description_id INT,
	operator_id INT,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--table product_type
CREATE TABLE product_type (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	name VARCHAR(64) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--table operator
CREATE TABLE operator (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	name VARCHAR(64) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--table product_description
CREATE TABLE product_description (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	description TEXT,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--table payment_method
CREATE TABLE payment_method (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	name VARCHAR(64),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--2.C.Create table transaction, transaction detail
--table transaction
CREATE TABLE transaction (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	user_id int,
	payment_method_id int,
	total_quantity int,
	total_price int,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--table transaction detail
CREATE TABLE transaction_detail (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	transaction_id int,
	product_id int,
	quantity int,
	price int,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


--3.Create tabel kurir dengan field id, name, created_at, updated_at.
CREATE TABLE kurir (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	name VARCHAR(64),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


--4.Tambahkan ongkos_dasar column di tabel kurir.
ALTER TABLE kurir ADD COLUMN ongkos_dasar int AFTER name;


--5.Rename tabel kurir menjadi shipping.
ALTER TABLE kurir RENAME shipping;

--6.Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
DROP TABLE shipping;


--7.Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
--7.A.1-to-1: payment method description.
--Untuk membuat relationship, table parent (payment_method_description) harus dibuat terlebih dahulu.
CREATE TABLE payment_method_description (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	description TEXT,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--tambah field/attribute pada tabel child (payment_method) untuk referensi relasi (payment_method_description_id)
ALTER TABLE payment_method ADD COLUMN payment_method_description_id int AFTER name;

--Buat constraint/relasi tabel
ALTER TABLE payment_method ADD FOREIGN KEY (payment_method_description_id) REFERENCES payment_method_description(id);

--7.B.1-to-many: user dengan alamat.
--ubah field/attribute alamat(address) user menjadi address_id int

--7.C.many-to-many: user dengan payment method menjadi user_payment_method_detail.
--buat table baru dengan membuat 2 constraint/relasi tabel
CREATE TABLE user_payment_method_detail (
	id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	payment_id INT,
	user_id INT,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (payment_id) REFERENCES payment_method(id),
	FOREIGN KEY (user_id) REFERENCES user(id)
);


--semua command untuk menghubungkan semua table
ALTER TABLE transaction ADD FOREIGN KEY (user_id) REFERENCES user(id);

ALTER TABLE transaction ADD FOREIGN KEY (payment_method_id) REFERENCES payment_method(id);

ALTER TABLE transaction_detail ADD FOREIGN KEY (transaction_id) REFERENCES transaction(id);

ALTER TABLE transaction_detail ADD FOREIGN KEY (product_id) REFERENCES product(id);

ALTER TABLE product ADD FOREIGN KEY (product_type_id) REFERENCES product_type(id);

ALTER TABLE product ADD FOREIGN KEY (product_description_id) REFERENCES product_description(id);

ALTER TABLE product ADD FOREIGN KEY (operator_id) REFERENCES operator(id);
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    