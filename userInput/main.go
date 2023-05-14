package main

import (
	"bufio"
	"fmt"
	"os"
)

var yazi string

func main() {
	//Scan fonksiyonunda dikkat edilmesi gereken nokta kullanıcı istediği kadar kelime girse bile programın ilk kelimeyi değer olarak alacağıdır. Scan() fonksiyonu boş giriş kabul etmez.
	//scan fmt paketi içerisinde bulunmaktadır.
	//fmt.Scan(&yazi)
	//fmt.Println(yazi)

	//Scanf() fonksiyonu Printf() fonksiyonu gibi format içerir.Bu fonksiyon ile kullanıcının girişini bölüp birkaç değişkene kaydedebiliriz.
	//Scanf(), boş giriş kabul eder.
	var k1, k2 string
	//adreslerini alarak input almayı unutma
	fmt.Scanf("%s %s", &k1, &k2)
	fmt.Println(k1)
	fmt.Println(k2)

	//Reader ile Satır Olarak Değer Alma : bir satır yazıyı giriş olarak alabilirsiniz.
	giris := bufio.NewReader(os.Stdin)
	yazi, _ := giris.ReadString('\n')
	fmt.Println(yazi)

}
