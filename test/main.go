package main

import "fmt"

func Merhaba(isim string) (cikti string) {
	cikti = "Merhaba " + isim
	return
}

func main() {
	fmt.Println(Merhaba("Ahmet"))
}
