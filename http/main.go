package main

import (
	"fmt"
	"net/http"
)

// dışarıdan w adında parametre alır.
// r:request nesnesinin pointerını alır.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:]) //url path'inde / 1.index olarak görünür. 1: ile / karakteri gelmez
}

func main() {
	//HandleFunc() fonksiyonu belirlediğimiz adrese girildiğinde hangi fonksiyonun çalıştırılacağınız belirliyor
	http.HandleFunc("/", handler)

	//ListenAndServe() fonksiyonu ise sunucunun ayağa kalkmasını ve istediğimiz bir porttan ulaşılmasını sağlıyor.
	http.ListenAndServe(":9000", nil)

	fmt.Println("Web sunucu")
}
