package lib

import (
	"fmt"
	"time"
)

// penerapan penggunaan keyword for - range - close pada channel
func SendMessage(msg chan string) {

	for i := 0; i < 20; i++ {
		now := time.Now()
		x, x1, x2 := time.Now().Clock()
		msg <- fmt.Sprintf("Hello, hari ini hari %s, jam %d:%d:%d:%d", time.Now().Weekday(), x, x1, x2, now.Nanosecond())
	}

	// digunakan untuk menutup atau menonaktifkan channel
	// setelah channel di close atau di nonaktifkan,
	// maka channel tidak bisa digunakan lagi
	close(msg)
}

func PrintMessage(ch chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}
