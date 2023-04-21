# Docker
Docker adalah aplikasi infrastructruce & deployment yang berbasis container

## Container 
Container bukanlah sebuah virtual machine tetapi proses (aplikasi) dengan sistem file isolasi. Virtual machine berjalan diatas Hypervisor/Virtual Machine manager dimana setiap Virtual Machine berjalan diatas OS nya masing-masing. Sedangkan Container berjalan diatas satu OS yang diatur oleh Container Manager.

## Docker Syntax
- FROM ~> Mengambil image dari docker registry
- RUN ~> Mengeksekusi Bash command saat build container
- ENV ~> Set sebuah variable dalam container
- ADD ~> meng-copy file dengan beberapa proses lainnya.
- COPY ~> copy file
- WORKDIR ~> mengatur working direktori atau direktori utama
- ENTRYPOINT ~> mengeksekusi command setelah container selesai di build
- CMD ~> mengeksekusi command tetapi bisa di timpa / overwrite
