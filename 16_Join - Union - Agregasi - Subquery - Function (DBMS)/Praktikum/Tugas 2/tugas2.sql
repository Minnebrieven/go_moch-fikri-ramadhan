--Join, Union, Sub query, Function
--1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT * FROM transaction WHERE user_id = 1 UNION SELECT * FROM transaction WHERE user_id = 2;

--2. Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(price) FROM transaction_detail WHERE transaction_id IN (SELECT id FROM transaction WHERE user_id = 1);

--3. Tampilkan total transaksi dengan product type 2.
SELECT COUNT(*) FROM transaction_detail td INNER JOIN product p ON td.product_id = p.id WHERE p.product_type_id = 2;

--4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT * FROM product p INNER JOIN product_type pt ON p.product_type_id = pt.id;

--5. Tampilkan semua field table transaction, field name table product dan field name table user.
--SELECT td.id, t.id AS transaction_id, u.name, p.name, t.payment_method_id FROM transaction_detail td INNER JOIN transaction t ON td.transaction_id = t.id INNER JOIN user u ON t.user_id = u.id INNER JOIN product p ON td.product_id = p.id;
SELECT t.*, p.*, u.* FROM transaction_detail td INNER JOIN transaction t ON td.transaction_id = t.id INNER JOIN user u ON t.user_id = u.id INNER JOIN product p ON td.product_id = p.id;

--6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER $$

CREATE FUNCTION deleteTransaction (idTransaction INT)
RETURNS INT DETERMINISTIC

BEGIN
SET FOREIGN_KEY_CHECKS=0;
DELETE FROM transaction WHERE id = idTransaction;
SET FOREIGN_KEY_CHECKS=1;
RETURN 1;
END $$

DELIMITER;

DELIMITER $$
CREATE TRIGGER deleteTransactionTrigger
BEFORE DELETE 
ON transaction FOR EACH ROW

BEGIN
DECLARE old_transaction_id INT;
SET old_transaction_id = OLD.id;
SET FOREIGN_KEY_CHECKS=0;
DELETE FROM transaction_detail WHERE transaction_id = old_transaction_id;
SET FOREIGN_KEY_CHECKS=1;
END $$

DELIMITER;

--7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER $$

CREATE FUNCTION deleteTransactionDetail (idTransactionDetail INT)
RETURNS INT DETERMINISTIC

BEGIN
SET FOREIGN_KEY_CHECKS=0;
DELETE FROM transaction_detail WHERE id = idTransactionDetail;
SET FOREIGN_KEY_CHECKS=1;
RETURN 1;
END $$

DELIMITER;

DELIMITER $$
CREATE TRIGGER subtractTransactionQuantity
AFTER DELETE 
ON transaction_detail FOR EACH ROW

BEGIN
UPDATE transaction SET total_quantity = total_quantity - OLD.quantity;
END $$

DELIMITER;


--8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT * FROM PRODUCT WHERE id NOT IN (SELECT product_id FROM transaction_detail);