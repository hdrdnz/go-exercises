package main

import (
	"fmt"
	"os"
)

func dosyaVarMı(isim string) bool {
	// os.Stat ile dosyanın bilgileri çekilir.
	bilgi, hata := os.Stat(isim)
	if hata != nil {
		return false
	}
	//Isdir() fonksiyonu dosya olup olmadığını kontrol eder.
	return !bilgi.IsDir()
}

func main() {
	if t := "dosya.txt"; dosyaVarMı(t) {
		fmt.Println(t, "bulunuyor")
	} else {
		fmt.Println(t, "bulunmuyor.")
	}
}
