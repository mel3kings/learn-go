package routines

import "fmt"

var naturals = make(chan int, 4)

func ReadFromBuffer() {
	for {
		x, ok := <-naturals
		if ok {
			fmt.Println("RECEIVING..", x)
		}
	}
}

func SendToBuffer() {
	for x := 0; x < 10; x++ {
		naturals <- x
	}
}

func main() {
	go SendToBuffer()
	go ReadFromBuffer()
}
