package main

import (
	"fmt"
)

type insan struct {
	k1, k2, k3 string
}

// Map ile  dizi içerisine yeni bir dizi ekliyorsunuz. Tabi bunu struct metodu ile yapıyoruz.
func main() {
	var a map[string]insan
	a = make(map[string]insan)
	a["erkekler"] = insan{
		"ali", "veli", "hüseyin",
	}
	a["kadınlar"] = insan{
		"ayşe", "alina", "damla",
	}
	fmt.Println(a)

}
