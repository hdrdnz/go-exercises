package main

import (
	"fmt"
	"net/http"
)

type Human struct {
	Fname string
	Lname string
	Age   int
}

// ServeHTTP implements http.Handler
func (Human) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("human bulunamadı")
}

func (h Human) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	h.Fname = "Havva Nur"
	h.Lname = "Durudeniz"
	h.Age = 22

	//formu parse etmek için:
	r.ParseForm()

	//sunucudan form bilgisini almak için:
	fmt.Println(r.Form)

	//URL'in path bilgisini almak için:
	fmt.Println("path", r.URL.Path)

	//fmt.Fprint(w, "<table><tr> <td><b>"+h.Fname+"</b></td><td><b>"+h.Lname+"</b></td>")

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var h Human

	err := http.ListenAndServe(":9000", h)
	checkError(err)

}
