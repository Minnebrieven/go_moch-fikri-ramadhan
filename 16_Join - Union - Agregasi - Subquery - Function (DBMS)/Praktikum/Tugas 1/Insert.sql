--Data Manipulation Language
--1.Insert
--1.A.Insert 5 operators pada table operator
INSERT INTO operator (name) VALUES ("Andy"), ("John"), ("Diane"), ("Rachel"), ("Mande");


--1.B.Insert 3 product type.
INSERT INTO product_type (name) VALUES ("Electronic"), ("Cosmetics"), ("Accessories");


--1.C.Insert 2 product dengan product type id = 1, dan operators id = 3.
INSERT INTO product (name, price, product_type_id, operator_id) VALUES 
("Realme C2", 1200000, 1, 3), 
("Go Pro Hero 12", 10000000, 1, 3);


--1.D.Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO product (name, price, product_type_id, operator_id) VALUES 
("Wardah Foundation", 150000, 2, 1), 
("Red Lipstick", 75000, 2, 1), 
("Fake eyelashes", 15000, 2, 1);


--1.E.Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO product (name, price, product_type_id, operator_id) VALUES 
("Silver Bracelet", 200000, 3, 4), 
("Sunglasses", 35000, 3, 4), 
("Silver Necklace", 300000, 3, 4);


--1.F. Insert product description pada setiap product.
--insert description terlebih dahulu untuk setiap product ke dalam product_description
INSERT INTO product_description (description) VALUES 
("Smartphone Brand Realme dengan model C2"), 
("Camera kecil yang tahan lama,tahan banting dan waterproof"), 
("Foundation dengan tekstur yang ringan untuk menyamarkan warna kulit tidak merata"), 
("Lipstick yang memiliki warna merah"), ("Bulu mata palsu untuk mempercantik bulu mata"), 
("Gelang yang dibuat dengan material silver"), ("Kacamata hitam guna melindungi mata dari sinar matahari"), 
("Kalung yang dibuat dengan material silver");

--ubah field/attribute product_description_id pada masing-masing product sesuaikan dengan id dari product_description
UPDATE `product` SET `product_description_id` = '1' WHERE `product`.`id` = 1;
UPDATE `product` SET `product_description_id` = '2' WHERE `product`.`id` = 2;
UPDATE `product` SET `product_description_id` = '3' WHERE `product`.`id` = 3;
UPDATE `product` SET `product_description_id` = '4' WHERE `product`.`id` = 4;
UPDATE `product` SET `product_description_id` = '5' WHERE `product`.`id` = 5;
UPDATE `product` SET `product_description_id` = '6' WHERE `product`.`id` = 6;
UPDATE `product` SET `product_description_id` = '7' WHERE `product`.`id` = 7;
UPDATE `product` SET `product_description_id` = '8' WHERE `product`.`id` = 8;


--1.G. Insert 3 payment methods.
INSERT INTO payment_method (name) VALUES ("Cash On Delivey"), ("Bank Transfer"), ("E-Wallet");


--1.H. Insert 5 user pada tabel user.
INSERT INTO user (name, birthdate, gender) VALUES 
("M Fikri Ramadhan", '2001-10-11', "L"), 
("Arbyansyah", '1998-6-24', "L"), 
("Devi Amalia", '2000-8-15', "P"), 
("M Hanif", '2002-12-1', "L"), 
("Sean Anthony", '2000-4-1', "L");


--1.I. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
--transaction user_id = 1
INSERT INTO transaction (user_id, payment_method_id, total_quantity, total_price) VALUES 
(1, 2, 3, 11235000), 
(1, 3, 3, 535000),
(1, 1, 3, 250000);

--transaction user_id = 3
INSERT INTO transaction (user_id, payment_method_id, total_quantity, total_price) VALUES 
(2, 2, 3, 240000),
(2, 3, 3, 575000),
(2, 3, 3, 10500000);

--transaction user_id = 4
INSERT INTO transaction (user_id, payment_method_id, total_quantity, total_price) VALUES 
(4, 2, 3, 10275000), 
(4, 3, 2, 300000),
(4, 1, 3, 1160000);



--1.J. Insert 3 product di masing-masing transaksi.
--transaction_id = 1
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(1, 1, 1, 1200000), (1, 2, 1, 10000000), (1, 7, 1, 35000);

--transaction_id = 2
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(2, 6, 1, 200000), (2, 7, 1, 35000), (2, 8, 1, 300000);

--transaction_id = 3
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(3, 5, 1, 15000), (3, 6, 1, 200000), (3, 7, 1, 35000);

--transaction_id = 4
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(4, 3, 1, 150000), (4, 4, 1, 75000), (4, 5, 1, 15000);

--transaction_id = 5
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(5, 4, 1, 75000), (5, 6, 1, 200000), (5, 8, 1, 300000);

--transaction_id = 6
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(6, 2, 1, 10000000), (6, 6, 1, 200000), (6, 8, 1, 300000);

--transaction_id = 7
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(7, 2, 1, 10000000), (7, 4, 1, 75000), (7, 6, 1, 200000);

--transaction_id = 8
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(8, 3, 1, 150000), (8, 5, 1, 15000), (8, 7, 1, 35000);

--transaction_id = 9
INSERT INTO transaction_detail (transaction_id, product_id, quantity, price) VALUES 
(9, 1, 1, 1200000), (9, 4, 1, 75000), (9, 7, 1, 35000);