package routines

import "fmt"

var naturals = make(chan int, 4) // buffered to handle 4 at a time

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
	close(naturals) // optional close channel, will no longer be able to send to it again.
}

func UnidirectionalChannel() {
	go SendToBuffer()
	//	go ReadFromBuffer()
	for x := range naturals { // another way of reading from channel
		fmt.Println("RECEIVING..", x)
	}

}
