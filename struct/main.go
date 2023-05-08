package main

import "fmt"

//type terimi ile yeni bir tür oluşturulur.
type kisi struct {
	isim    string
	soyisim string
	yas     int
}

//struct özel fonksiyon
//i: kisi struct içerisindeki değişkenlere ulaşmak için bu değişken kullanılır.
func (i kisi) tanit() {
	fmt.Printf("Merhaba bne %s. %d yasındayım.", i.isim, i.yas)
}

func main() {
	kisi1 := kisi{"ahmet", "erdem ", 5}
	kisi2 := kisi{}
	kisi2.isim = "ayşe"
	kisi2.soyisim = "eslem"
	kisi2.yas = 50
	fmt.Println(kisi2)
	kisi1.tanit()

	//anonim struct
	ögrenci := struct {
		ad, bölüm string
	}{"murat", "ceng"}

	fmt.Println(ögrenci)
}
