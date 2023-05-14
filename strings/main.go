package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var eposta string
	fmt.Print("e posta adresinizi giriniz:")
	fmt.Scan(&eposta)
	//Contains():string içerisinde istenilen ifade var mı diye kontrol yapar.
	if strings.Contains(eposta, ".com") {
		fmt.Println("doğru eposta")
	} else {
		fmt.Println("yanlış eposta")
	}
	//Count() fonksiyonu ile bir string değerin içinde istediğimiz bir string değerin kaç tane olduğunu öğrenebiliriz
	fmt.Println(strings.Count(eposta, "h"))

	//Index() fonksiyonu ile bir string değerin içindeki istediğimiz bir string değerin kaçıncı sırada yani index’te olduğunu öğrenebiliriz
	fmt.Println(strings.Index("istemek", "e"))

	//LastIndex() fonksiyonu ile bir string değerin içinde istediğimiz bir string değerin sırasını Index() fonksiyonunun tersine sağdan sola doğru kontrol eder.

	//Title() fonksiyonu ile içerisine küçük harflerle string türünde değer girdiğimizde baş harfleri büyük harf yapan bir fonksiyondur.

	//ToUpper() fonksiyonu içerisine girilen string değerin tüm harflerini büyük harf yapar.
	//ToLower() fonksiyonu içerisine girilen string değerin tüm harflerini küçük harf yapar.

	fmt.Println(strings.ToUpper("alime"))
	//ToUpper() fonksiyonu ile string değeri büyük harf yaptığımız zaman Türkçe karakter sıkıntısı yaşarız. Örnek olarak “i” harfi büyüyünce “I” harfi olur. Bunun önüne ToUpperSpecial() fonksiyonu ile geçebiliriz.

	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "alime"))

	//ToUpperSpecial() fonksiyonu ile aynı seçilde çalışır ;fakat harfleri belirlediğiniz karakter kodlamasına göre küçültür
}
