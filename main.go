package main

import (
	"fmt"
	"runtime"
)

func main() {

	if r := runtime.GOOS; r == "windows" {
		fmt.Println("Windows için yönetici olarak çalıştırın.")
	} else if r == "linux" {
		fmt.Println("Linux için sudo komutu ile çalıştırın.")
	} else {
		fmt.Println("Geçersiz işletim sistemi!")
	}
}
