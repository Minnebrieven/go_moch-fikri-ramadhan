--4. Delete
--4.A. Delete data pada tabel product dengan id = 1.

--untuk disable foreign key check agar bisa menghapus walaupun terdapat relasi
SET FOREIGN_KEY_CHECKS=0;
DELETE FROM product WHERE id = 1;
--enable kembali
SET FOREIGN_KEY_CHECKS=1;


--4.B. Delete pada pada tabel product dengan product type id 1.

--untuk disable foreign key check agar bisa menghapus walaupun terdapat relasi
SET FOREIGN_KEY_CHECKS=0;
DELETE FROM product WHERE product_type_id = 1;
--enable kembali
SET FOREIGN_KEY_CHECKS=1;