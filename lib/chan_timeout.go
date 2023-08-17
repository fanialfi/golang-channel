package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func SendData(ch chan<- int) {
	for i := 0; true; i++ {
		// mengirimnakn data pada interval waktu tertentu, dimana durasinya adalah acak / random
		ch <- i

		// time.Sleep() akan menjeda ke perulangan selanjutnya dengan durasi yang acak
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func RetriveData(ch <-chan int) {
loop:
	for {
		select {

		// ketika sebuah data pada channel disimpan dalam dua variabel, maka variabel kedua berisi kebenaran datanya
		// apakah datanya terkirim atau tidak
		case data, open := <-ch:
			fmt.Printf("receive data \"%v\" bernilai %t\n", data, open)
		case <-time.After(time.Second * 5):
			fmt.Printf("timeout...\nno activate after 5 seconds\n")
			break loop
			// close(ch) // cannot close channel receive only
			// function close hanya bisa digunakan untuk channel direction bertipe sender dan booth
		}
	}
}
