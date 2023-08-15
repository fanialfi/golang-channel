package lib

import "fmt"

func DisplayLoop(ch chan int) {
	for {
		fmt.Printf("receive message : %d\n", <-ch)
	}
}
