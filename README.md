# Channel

channel digunakan untuk menghubungkan goruntine satu dengan goruntine yang lain.
Dalam komunikasi-nya channel digunakna sebagai pengirim pada goruntine, dan penerima pada goruntine lain.
Proses pengiriman dan penerimaan data pada channel bersifat **blocking** atau _synchronous_.

![proses pengiriman dan penerimaan data pada channel][channel]

channel merupakan sebuah variabel, dibuat dengan menggunakan keyword `make` dan `chan`, variabel channel bertugas sebagai pengirim atau penerima sebuah data.

contoh penerapan channel :

```go
package main

import (
  "fmt"
  "runtime"
)

func main(){
  runtime.GOMAXPROCS(4)

  messages := make(chan string)

  sayHelloTo := func(to string){
    data := fmt.Sprintf("Hello %s", to)

    messages <- data
  }

  // proses di bawan ini bersifat asynchronous
  go sayHelloTo("fani")
  go sayHelloTo("alfi")
  go sayHelloTo("fanialfi")


  // proses di bawan ini bersifat blocking
  message1 := <- messages
  fmt.Println(message1)

  message2 := <- messages
  fmt.Println(message2)

  message3 := <- messages
  fmt.Println(message3)
}
```

pada kode di atas variabel `messages` dideklarasikan sebagai channel string.
cara pembuatanya dengan cara menggunakan keyword `make` dengan isi keyword `chan` lalu diikuti tipe data channel yang diinginkan, pada contoh di atas variabel `messages` bertipe data channel string.

ada juga closure `sayHelloTo` yang digunakna untuk membuat pesan string, yang kemudian dikirim via channel `messages <- data`.
Tanda `<-` jika ditulis di sebelah kiri nama variabel, berarti sedang berlangsung proses pengiriman data dari variabel yang ada di kanan lewat channel yang berada di kiri.

function `sayHelloTo` dieksekusi tiga kali sebagai goruntine yang berbeda, yang membuat prosesnya menjaid asynchronous.

dari ketiga goruntine tersebut yang paling awal selesai mengirimkan data dulu, datanya kemudian diterima oleh variabel `message1`.
Tanda `<-` jika dituliskan sebelah kiri adalah channel, menandakan proses penerimaan data dari channel yang dikanan untuk disimpan di variabel yang dikiri.

```go
  message1 := <- messages
  fmt.Println(message1)

  message2 := <- messages
  fmt.Println(message2)

  message3 := <- messages
  fmt.Println(message3)
```

penerimaan data dari channel bersifat blocking, artinya statement `message1 := <- messages` hingga setelahnya tidak akan di eksekusi sebelum ada data yang dikirim lewat channel.

selain itu variabel channel juga bisa di pass sebagai parameter pada function.
Cukup tambahkan keyword `chan` pada deklarasi parameter agar operasi pass channel variabel bisa dilakukan.

## buffered channel

proses kirim data pada channel secara default dilakukan secara synchronous (**blocking**) atau tidak di **buffer** di memori.
Ketika terjadi proses kirim-terima data pada sebuah goruntine, maka harus ada goruntine lain yang menerima data pada channel yang sama.

Pada buffered channel dilakukan sedikit berbeda, pada channel jenis ini, ditentukan juga jumplah buffer nya,
angka tersebut menjadi penentu jumplah data yang dapat diterima secara bersamaan selama jumplah data tidak melebihi jumplah buffer yang ditentukan.

Jika jumplah data sudah melewati batas buffer, maka pengiriman selanjutnya hanya bisa dilakukan ketika satu data yang sudah terkirim sudah diambil oleh channel goruntine yang lain, sehingga ada slot yang kosong.

Proses pengiriman data pada buffered channel bersifat asynchronous, namun ketika jumplah data yang dikirim sudah melebihi batas maksimum buffer, maka proses selanjutnya bersifat synchronous.

![proses pengiriman bufered channel][buffer]

Untuk penerapan buffered channel, sama seperti penerapan channel pada umumnya, perbedaannya pada saat deklarasi saja.
saat pembuatan buffered channel menggunakan keyword `make` parameter kedua diisi jumplah / ukuran buffer yang dapat ditampung.

Nilai buffer pada channel dimulai dari 0, maka jika nilainya 3 berarti jumplah maksimal adalah 4.

```go
package main

import (
  "fmt"
  "runtime"
)

func main(){
  runtime.GOMAXPROCS(4)

  message := make(chan string, 3)
	go func() {
		for {
			i := <-message
			fmt.Printf("receive data : %s\n", i)
		}
	}()

	
	for i := 1; i <= 5; i++ {
		fmt.Printf("sending data-%d\n", i)
		data := fmt.Sprintf("sending data-%d\n", i)
		message <- data
	}
}
```

[channel]: ./img/channel.png
[buffer]: ./img/channel-buffer.png