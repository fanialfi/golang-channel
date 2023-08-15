package lib

// function untuk mencari nilai rata rata
func GetAvg(numbers []int, ch chan float64) {
	sum := 0

	for _, elm := range numbers {
		sum += elm
	}

	ch <- float64(sum) / float64(len(numbers))
}

// function untuk mencari nilai tertinggi
// nilai tertinggi akan dikirim via channel
func GetMax(numbers []int, ch chan int) {
	max := numbers[0]

	for _, elm := range numbers {
		if max < elm {
			max = elm
		}
	}

	ch <- max
}
