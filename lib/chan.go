package lib

import "fmt"

func SayHelloTo(name string, ch chan string) {
	data := fmt.Sprintf("Hello %s", name)

	ch <- data
}

// penggunaan channel pada goruntine
