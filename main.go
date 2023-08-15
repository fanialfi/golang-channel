package main

import (
	"fmt"
	"runtime"
	"time"
)

// function untuk mencari nilai rata rata
func getAvg(numbers []int, ch chan float64) {
	sum := 0

	for _, elm := range numbers {
		sum += elm
	}

	ch <- float64(sum) / float64(len(numbers))
}

// function untuk mencari nilai tertinggi
// nilai tertinggi akan dikirim via channel
func getMax(numbers []int, ch chan int) {
	max := numbers[0]

	for _, elm := range numbers {
		if max < elm {
			max = elm
		}
	}

	ch <- max
}

func sayHelloTo(name string, ch chan string) {
	data := fmt.Sprintf("Hello %s", name)

	ch <- data
}

func displayLoop(ch chan int) {
	for {
		fmt.Printf("receive message : %d\n", <-ch)
	}
}

// penerapan penggunaan keyword for - range - close pada channel
func sendMessage(msg chan string) {

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

func printMessage(ch chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	awal := time.Now()
	runtime.GOMAXPROCS(4)

	messages := make(chan string, 2)

	go sayHelloTo("fani", messages)
	go sayHelloTo("alfi", messages)
	go sayHelloTo("fanialfi", messages)

	// contoh penggunaan buffered channel
	message := make(chan int, 2)
	go displayLoop(message)

	func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("sending data-%d\n", i)
			message <- i
		}
	}()

	ch1 := make(chan float64, 4)
	ch2 := make(chan int, 4)

	numbers := []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)

	go getAvg(numbers, ch1)
	go getMax(numbers, ch2)

	// penggunaan keyword select pada channel
	for i := 0; i < 5; i++ {
		micro := time.Now().UnixMicro()
		select {
		case msg1 := <-messages:
			fmt.Printf("pesan ke-1, %s, selisih waktu : %d microsecond\n", msg1, time.Now().UnixMicro()-micro)
		case msg2 := <-messages:
			fmt.Printf("pesan ke-2, %s, selisih waktu : %d microsecond\n", msg2, time.Now().UnixMicro()-micro)
		case msg3 := <-messages:
			fmt.Printf("pesan ke-3, %s, selisih waktu : %d microsecond\n", msg3, time.Now().UnixMicro()-micro)
		case avg := <-ch1:
			fmt.Printf("rata rata\t : %.2f\n", avg)
		case max := <-ch2:
			fmt.Printf("max\t\t : %d\n", max)
		}
	}

	// penerapan for range close pada channel
	forRangeChan := make(chan string, 2)
	go sendMessage(forRangeChan)
	printMessage(forRangeChan)

	akhir := time.Now()
	fmt.Printf("\nprogram berjalan selama : %v\n", akhir.Sub(awal))
}
