package main

import "fmt"

func Tamİsim(isim *string, soyisim *string) {
	if isim == nil {
		panic("isim nil olamaz.")
	}
	if soyisim == nil {
		panic("soyisim nil olamaz.")
	}
	fmt.Printf("%s %s\n", *isim, *soyisim)
	fmt.Println("tamİsim fonksiyonu tamamlandı")

}

func main() {
	ad := "ahmet"
	soyisim := "erdem"
	defer func() {
		// r değişkenini recover türünden tanımladık. r ye bir panik mesajı gelirse yani nil den farklı olursa bu if bloğu çalışır.
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	Tamİsim(&ad, nil)
	Tamİsim(nil, &soyisim)
}

//eğer istenmeyen bir durum var ise ana metot durur. ana metodun durmasını engellemek içn recover() kullanabilirsin.
