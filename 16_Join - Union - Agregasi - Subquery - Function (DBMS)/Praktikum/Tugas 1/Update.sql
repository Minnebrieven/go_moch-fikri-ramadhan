--3.Update
--3.A. Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE product SET name = "product dummy" WHERE id = 1;

--3.B. Update qty = 3 pada transaction detail dengan product id = 1.
UPDATE transaction_detail SET quantity = 3 WHERE product_id = 1;
