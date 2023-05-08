package main

import "fmt"

var a [3]string

//boş bir dilimin varsayılan değeri null'dur.
var b []int

func main() {
	nums := [6]int{4, 8, 7, 1, 5, 1}
	//dinamik büyüklükte diziler oluşturmak için make fonksiyonu kullanılır.
	c := make([]int, 7)
	a[0] = "ahmet"
	a[1] = "ayşe"
	a[2] = "selim"
	//uzunluğunu vermediğimiz dizilere ekleme yaparken append metodu kullanılır.
	b = append(b, 58)
	b = append(b, 87)
	//Range, üzerinde kullanıldığı diziyi for döngüsü ile tekrarlayabilir.
	for x, y := range a {
		fmt.Printf("%d. indeks = %s\n", x, y)
	}
	fmt.Println(len(a))
	fmt.Println(nums)
	fmt.Println(nums[:4])
	fmt.Println(b)
	fmt.Println(c)
}
