package main

import "fmt"

//Sprintf fonksiyonu fmt paketine dahil bir fonksiyondur. Bu fonksiyon değişkenlere formatlı atama yapmamıza yardımcı olur.

func main() {
	ad := "murat"
	adTip := fmt.Sprintf("bu isim %T tipindedir", ad)
	fmt.Println(adTip)
}
