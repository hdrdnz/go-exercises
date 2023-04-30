package main

import "fmt"

func main() {
	a := 8
	//ekle(a) //bu şekilde yaparsan a da herhangi bir değişiklik olmaz. yine 8 değerini basacaktır. a nın değerini değiştirmek için adresini kullanmamız gerekiyor. bu yüzden burada pointer yapısını kullanarak adresine ulaşabiliriz.

	ekle(&a) //burada & işareti kullanarak ekle fonksiyonuna a nın adresini gönderdik.
	fmt.Println(a)

}

//gelen adresin değerini almak için de "*" operatörü kullanılır.
func ekle(x *int) {
	*x += 5
}
