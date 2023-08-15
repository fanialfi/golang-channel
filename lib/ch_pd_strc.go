package lib

import (
	"fmt"
	"time"
)

// penggunaan channel pada struct

type Person struct {
	Name string
	Data chan string
}

func (p Person) SendMessage(ch chan string) {
	for i := 0; i < 20; i++ {
		x, x1, x2 := time.Now().Clock()
		now := time.Now()
		ch <- fmt.Sprintf("hallo %s, hari ini hari %s, jam %d:%d:%d:%v", p.Name, time.Now().Weekday(), x, x1, x2, now.Nanosecond())
	}
	close(ch)
}

func (p Person) PrintMessage(ch chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}
