# Version Control and Branch Management (Git)
Apa itu versioning? Versioning biasanya dimanfaatkan untuk mengatur versi dari suatu source code program. Terdapat berbagai macam alat(tools) yang dapat digunakan untuk mengatur versi dari suatu source program diantaranya :
1. Version Control System (VCS)
2. Soruce Code Manager (SCM)
3. Revision Control System (RCS)


### Jenis-jenis Version Control System (VCS)
1. Single User


Jenis ini merupakan yang paling sederhana dan hanya terdapat pada satu local komputer (Localized) saja

2. Centralized


Pada system ini terdapat satu system terpusat di internet ataupun local. Para developer sangat bergantung pada satu system pusat ini, jika system pusat tersebut mati maka para developer tidak dapat mengembangkan programnya.

3. Distributed


Jenis system ini paling banyak digunakan dan efektif semenjak tahun 2005 salah satunya Git. System ini tidak hanya menyimpan source code di system pusat tetapi juga menyimpan di local. Pada system pusat disana lah tempat memanage source code, para developer dapat mengembangkan programnya di komputer masing-masing (local) tetapi syncronized (terhubung ke system pusat)

### GIT
Git adalah salah satu Version Control System yang paling banyak dan populer digunakan oleh para developer dalam mengembangkan software bersama-sama. Tidak hanya untuk mengembangkan software secara bersama-sama Git dapat tracking aktifitas yang terjadi dalam pengembangan software sehingga dapat mempermudah kolaborasi dengan developer lainnya.


### Command yang sering digunakan dalam Git
Sebelum mengunakan Git ini diharuskan membuat repository terlebih dahulu dimana repository tersebut yang akan menampung source code yang akan dikembangkan.


Untuk melanjutkan ketahap selanjutnya kita setting config dari git
```git
git config --global user.name "Masukan Username dari akun Git disini"
git config --global user.email "Masukan Email dari akun Git disini"
```

Setelah selesai setting config, maka selanjutnya menginisiasi git di file yang akan disimpan di repository
```git
git init
git remote add <nama_remote> <url_remote_repo>
```

Ketika tahap pengembangan selesai maka source program dapat diupload / dipush ke repository
```git
git add .                               //untuk menambah file working directory ke staging area
git commit -m "Masukan keterangan"      //untuk commit staging ke repository
git push origin main                    //untuk push ke repository
```

