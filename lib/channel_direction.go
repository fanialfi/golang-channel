package lib

import (
	"fmt"
	"time"
)

// parameter ch hanya bisa digunakan untuk mengirimkan data
func Send(ch chan<- string) {
	x, x1, x2 := time.Now().Clock()
	now := time.Now()
	ch <- fmt.Sprintf("sekarang jam %d:%d:%d:%d", x, x1, x2, now.Nanosecond())
}

// parameter ch hanya bisa digunakan untuk menerima data
func Receive(ch <-chan string) {
	fmt.Println(<-ch)
}

// parameter ch bisa digunakan untuk mengirim maupun menerima data
func Booth(ch chan string) {
	now := time.Now()
	x, x1, x2 := now.Clock()

	ch <- fmt.Sprintf("jam %d:%d:%d:%d", x, x1, x2, now.Nanosecond())

	fmt.Println("booth :", <-ch)
}
