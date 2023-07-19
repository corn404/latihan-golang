# Go Module

## Membuat Module

* untuk membuat module baru, kita bisa menggunakan perintah ```go mod init nama-module```
* Go-Lang akan secara otomatis membuat file ```go.mod``` yang berisikan informasi nama module dan juga versi Go-Lang yang kita gunakan

## Rilis Module

* Go-Lang terinteraksi baik dengan GIT
* Untuk merilis module, kita hanya perlu membuat Tag di GIT

## Menambah Dependency

* ```go get nama-module```

## Upgrade Module

* Untuk melakukan upgrade module, kita hanya cukup membuat tag baru di GIT

## Upgrade Dependency

* untuk upgrade Dependency ke versi terbaru, kita bisa mengubah isi ```go.mod```, lalu mengubah tag nya menjadi terbaru
* untuk mendownload versi terbaru, gunakan perintah ```go get```

## Major Upgrade

* Major upgrade biasanya tejadi diakrenakan ada perubahan pada isi kode program kita, sehingga membuatnya tidak backward compatible
* sebaiknya hal ini sebisa mungkin dihindari
* namun jika tidak bisa dihindari, strategy tebaik adalah merubah nama module
