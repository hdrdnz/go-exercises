package main

import "fmt"

type dynamic any

func main() {
	var x dynamic
	x = 13
	x = "hello"
	fmt.Printf("%T:%v\n", x, x)
}

//Yukarıdaki örnekte görüldüğü üzere x değişkenimize hangi tipte atama yaparsak, x değerin tipine dönüşüyor.
