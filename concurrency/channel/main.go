package main

import "fmt"

/*
使用了大量的 goroutine，
中間需要交換資料，這時候就可以使用 Channel 來進行溝通
*/
func main() {
	foo := addByShareCommunicate(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}

func addByShareCommunicate(n int) []int {
	var ints []int
	channel := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(channel chan<- int, order int) {
			channel <- order
		}(channel, i)
	}

	for i := range channel {
		ints = append(ints, i)

		if len(ints) == n {
			break
		}
	}

	close(channel)

	return ints
}
