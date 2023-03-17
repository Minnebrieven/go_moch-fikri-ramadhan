--2.Select
--2.A. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT name FROM user WHERE gender="L";

--2.B.Tampilkan product dengan id = 3.
--SELECT * FROM product WHERE id=3;
SELECT name FROM product WHERE id=3;

--2.C.Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM user WHERE name LIKE "%a%" AND created_at > TIMESTAMP(DATE_SUB(NOW(), INTERVAL 7 day));

--2.D.Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(gender) FROM user WHERE gender = "P";

--2.E.Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM user ORDER BY name ASC;

--2.F.Tampilkan 5 data pada data product
SELECT * FROM product LIMIT 5;

