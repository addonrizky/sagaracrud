This is Sagara API CRUD...

for quick run, to see working API, just do:
docker-compose up --build

then import "sagaracrud.postman_collection.json" into your postman, to simulate the API

things to notice: 
- API menggunakan JWT untk metode autentikasi-nya
- API tidak bisa di-call bila tidak menyertakan token JWT yg didapat ketika proses login 
- utk proses upload file, endpoint CREATE maupun UPDATE menerima string base64. Diasumsikan client telah terlebih dahulu mengencode file menjadi bentuk string base64 sebelum dikirim ke "sagara API CRUD"
- database utk kebutuhan API telah include dalam docker-compose, sehingga tidak perlu repot utk men-setup database untuk bisa mencoba working API
- API dikembangkan dengan pendekatan clean-code-architecture, harapannya agar lebih modular dan memudahkan dalam proses unit testing ke depannya
- karena keterbatasan waktu, unit test tidak diterapkan pada semua class, hanya class yang berada di package repository dan usecase saja yg sudah diberikan unit test
- dockerfile, docker-compose dan postman.collection sudah tersertakan pada repository. jadi feel free utk melihat-lihat




regards,
Rizky Ramadhan
