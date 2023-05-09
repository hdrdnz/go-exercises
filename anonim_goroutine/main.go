package main

import (
	"fmt"
)

func main() {
	kanal := make(chan string)
	go func() {
		//time.Sleep(2 * time.Second)
		fmt.Println("anonim goroutine")
		kanal <- "kanal bitti"

	}()
	fmt.Println("main fonksiyonu")
	//chan tipindeki bir değeri yazdırmak için başına <- işaretini koy.
	fmt.Println(<-kanal)

}
