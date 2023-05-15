package main

import (
	"fmt"
	"os"
)

func main() {
	//.Open() fonksiyonu ile dosyayı açıyoruz.
	dosya, err := os.Open("veri.xml")
	if err != nil {
		panic(err)
	}
	//dosyayı kapatmayı unutma.
	defer dosya.Close()

	fmt.Println("dosya başarıyla açıldı.")
}
