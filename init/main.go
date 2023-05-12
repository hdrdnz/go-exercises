package main

import "fmt"

//Golang’te bir uygulama çalışırken genelde çalışan ilk fonksiyon main() fonksiyonu oluyor. Bazen programın açılışında ayarlamamız gereken ön durumlar oluşuyor.bu durumda ise init  fonksiyonu devreye girer.
//Örnek olarak init() fonksiyonunda veritabanı bağlantımızı, kayıt defteri işlemlerimizi veya sadece bir kez yapmamız gereken işleri yapabiliriz.
//
var degisken string

func init() {
	degisken = "merhaba go"
}
func main() {
	fmt.Println(degisken)
}

//normalde init olmasaydı bu işlemi pointer ile gerçekleştirecektik.
