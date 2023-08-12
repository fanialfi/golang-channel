package main

import (
	"fmt"
	"runtime"
	"time"
)

// channel sebagai tipe data parameter
//
// passing data bertipe channel lewat parameter bersifat "pass by reference", yang di transmisikan adalah pointer datanya, bukan nilai datanya
func printMessage(msg chan string) {
	fmt.Println(<-msg)
}

func main() {
	runtime.GOMAXPROCS(4)

	messages := make(chan string, 5)

	sayHelloTo := func(name string) {
		data := fmt.Sprintf("Hello %s", name)

		messages <- data
	}

	go sayHelloTo("fani")
	go sayHelloTo("alfi")
	go sayHelloTo("fanialfi")

	go func() {
		message1 := <-messages
		fmt.Println(message1)

		message2 := <-messages
		fmt.Println(message2)

		message3 := <-messages
		fmt.Println(message3)
	}()

	// menggunakan channel sebagai tipe data parameter
	for _, each := range []string{"fani", "alfi", "fanialfi"} {
		micro := time.Now().UnixMicro()

		// menjalankan gorountine di anonymous function
		// anonymous function dibawah ini bersifat asynchronous
		go func(who string) {
			data := fmt.Sprintf("Selamat Pagi %s\t, selisih waktu : %d microsecond", who, time.Now().UnixMicro()-micro)

			messages <- data
		}(each)
	}

	{
		// ketika terjadi proses kirim data via channel dari sebuah goruntine,
		// maka harus ada goruntine lain yang menerima data dari channel yang sama,
		// dengan proses serah terima yang bersifat blocking
		//
		// Maksudnya, baris kode setelah kode pengiriman dan penerimaan data tidak akan di proses sebelum proses serah terima itu sendiri selesai
		// seperti contoh berikut ini :
		data := make(chan string)
		go func(data chan string) {
			fmt.Println(<-data)
		}(data)

		go func() {
			msg := fmt.Sprintf("sekarang hari %v", time.Now().Weekday())

			data <- msg
		}()
		// var pesan string
		// fmt.Scanln(&pesan)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}

	// contoh penggunaan buffered channel
	message := make(chan string, 3)
	go func() {
		for {
			i := <-message
			fmt.Printf("receive data : %s\n", i)
		}
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("sending data-%d\n", i)
			data := fmt.Sprintf("sending data-%d\n", i)
			message <- data
		}
	}()

	// proses dibawah digunakan sebagai blocking proses di atas,
	// karena proses di atas dilakukan secara asynchronous semua
	var str string
	fmt.Scanln(&str)
}
