<!-- awal readme.md -->
# Learn Golang

Selamat datang di repositori **Learn Golang**! Proyek ini dirancang untuk membantu Anda memulai dengan bahasa pemrograman Go. Baik Anda baru dalam pemrograman atau hanya baru dalam Go, panduan ini akan membantu Anda menyiapkan lingkungan Anda dan menjalankan program Go pertama Anda.

## Daftar Isi

- [Pendahuluan](#pendahuluan)
- [Prasyarat](#prasyarat)
- [Instalasi](#instalasi)
- [Menjalankan Program Go Pertama Anda](#menjalankan-program-go-pertama-anda)
- [Berpartisipasi](#berpartisipasi)
- [Lisensi](#lisensi)

## Pendahuluan

Go, juga dikenal sebagai Golang, adalah bahasa pemrograman open-source yang dirancang untuk kesederhanaan, efisiensi, dan keandalan. Repositori ini berisi contoh, latihan, dan penjelasan untuk membantu Anda belajar Go dengan efektif.

## Prasyarat

Sebelum Anda memulai, pastikan Anda telah menginstal hal berikut di komputer Anda:

- [Git](https://git-scm.com/)
- Editor kode atau IDE (seperti [Visual Studio Code](https://code.visualstudio.com/))

## Instalasi

1. **Instal Go**

   Ikuti instruksi di [situs resmi Go](https://golang.org/doc/install) untuk mengunduh dan menginstal Go sesuai sistem operasi Anda.

   Verifikasi instalasi dengan menjalankan perintah berikut di terminal Anda:

   ```bash
   go version
   ```

2. **Buat project Go**

   Buat direktori baru untuk proyek Anda dan masuk ke dalamnya:

   ```bash
   mkdir project-go
   cd project-go
   ```

   Setelah itu, Anda bisa mulai membuat file Go pertama Anda dengan menggunakan editor pilihan Anda.

3. **Inisialisasi Go Module**

   Go 1.11 atau yang lebih baru mendukung manajemen dependensi menggunakan Go Modules. Untuk menginisialisasi Go Module di dalam proyek Anda, jalankan perintah berikut di dalam direktori proyek:

   ```bash
     go mod init nama-proyek-anda
      ```

   Ganti nama-proyek-anda dengan nama yang sesuai untuk proyek Anda. Contohnya, jika proyek Anda disebut learn-golang, maka jalankan:

      ```bash
     go mod init learn-golang
      ```

   Perintah ini akan membuat file go.mod yang mencatat dependensi-proyek Go Anda.

## Menjalankan Program Go Pertama Anda
Untuk menjalankan program Go pertama Anda, buatlah file dengan ekstensi <code>.go</code> dan tulislah kode program Go di dalamnya. Misalnya, buat file hello.go dengan isi sebagai berikut:

   ```go
   package main

   import "fmt"

   func main() {
      fmt.Println("Hello, Go!")
   }
   ```

   Simpan file tersebut, dan jalankan dari terminal dengan perintah:

   ```bash
   go run hello.go
   ```

## Berpartisipasi

   Kami sangat mengundang kontribusi Anda untuk pengembangan repositori ini. Jika Anda menemukan bug, memiliki ide perbaikan, atau ingin menambahkan fitur baru, silakan buka issue atau kirim pull request.

## Lisensi

   Proyek ini dilisensikan di bawah lisensi MIT, yang berarti Anda dapat menggunakan, menyalin, memodifikasi, menggabungkan, menerbitkan, mendistribusikan, mensublisensikan, dan/atau menjual salinan perangkat lunak ini, asalkan mencantumkan pemberitahuan lisensi ini.
<!-- akhir readme.md -->