package main

import (
	"fmt"
	"io/ioutil"
)

//ioutil paketi standart Golang paketleri içerisinde gelir ve dosya işlemleri yapabilmemiz için bize fonksiyonlar sağlar.

func main() {
	//okunacak dosyayı belirtiyoruz.
	dosya, hata := ioutil.ReadFile("dosya.txt")

	if hata != nil {
		panic(hata)
	}

	//dosyanın içeriği ekrana bastırılır.
	//Okuma işlemi byte tipinde yapıldığı için string() fonksiyonu ile byte tipini string tipine dönüştürüyoruz.
	fmt.Println(string(dosya))
	//yazmak istediğimiz ifade belirtilir.
	veri := []byte("golang")
	err := ioutil.WriteFile("dosya.txt", veri, 0644)
	deger, err2 := ioutil.Discard.Write(veri)
	if err2 != nil {
		panic(err2)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(deger)
}
