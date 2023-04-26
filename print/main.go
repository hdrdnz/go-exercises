package main

import "fmt"

//değişken isimlendirme de camel case isimlendirme kullanılır.
//kısaltmaların hepsi büyük harfle yazılır.
//main dışında bu şekilde gösterim yapamazsın.
//a :=45
//var a int=45 bu şekilde değişken tanımlama yapabilirsin.

func main() {
	name := "murat"
	x := 100
	y := 5
	//benim adım murat
	fmt.Println("Benim adım", name)
	//benim adımmurat
	fmt.Print("benim adım", name)
	fmt.Println("")
	//printf içerisinde özel karakterler kullanmalısın.
	// %T =tipini yazdırır.
	//%v değerini yazdırır.
	fmt.Printf("benim adım %v", name)

	// %b : ikilik tabandaki değeri
	fmt.Printf("%b", x)
	fmt.Println("")
	fmt.Printf("%d", y)

	Name, age := "ayşe", 35
	fmt.Print("benim adım:", Name, "yaşım", age)

}
