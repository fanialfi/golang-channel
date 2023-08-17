package main

import (
	"fmt"
	"golang/channel/lib"
	"runtime"
	"time"
)

func main() {
	awal := time.Now()
	runtime.GOMAXPROCS(4)

	messages := make(chan string, 2)

	go lib.SayHelloTo("fani", messages)
	go lib.SayHelloTo("alfi", messages)
	go lib.SayHelloTo("fanialfi", messages)

	// contoh penggunaan buffered channel
	message := make(chan int, 2)
	go lib.DisplayLoop(message)

	func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("sending data-%d\n", i)
			message <- i
		}
	}()

	// channel sebagai parameter
	ch1 := make(chan float64, 4)
	ch2 := make(chan int, 4)

	numbers := []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)

	go lib.GetAvg(numbers, ch1)
	go lib.GetMax(numbers, ch2)

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
	go lib.SendMessage(forRangeChan)
	lib.PrintMessage(forRangeChan)

	// penggunaan channel pada struct
	fmt.Println()
	alpi := lib.Person{Name: "alfi", Data: make(chan string)}
	go alpi.SendMessage(alpi.Data)
	alpi.PrintMessage(alpi.Data)

	// contoh penggunaan channel direction
	fmt.Println()
	chanDirection := make(chan string, 3)
	go lib.Send(chanDirection)
	go lib.Booth(chanDirection)
	lib.Receive(chanDirection)

	// digunakan untuk blocking proses diatas
	// pada penggunaan channel direction
	var str string
	fmt.Scanln(&str)

	// implementasi channel timeout
	message = make(chan int, 2)
	go lib.SendData(message)
	lib.RetriveData(message)

	akhir := time.Now()
	fmt.Printf("\nprogram berjalan selama : %v\n", akhir.Sub(awal))
}
