//  channels,goroutinlerin birbirleriyle iletişime geçmesini sağlar
//  channels,bir goroutine tarafından gönderilen değerin main fonksiyonunda kullanılmadan önce bize gelmesini sağlar.
// !!goroutine return geri döndüren fonksiyonlarda kullanılmaz.ama channel ifade kullanılarak channeller ile return ifade bir kanala gönderilir.
//channeller bloklama işlemi yapar.

package main

import (
	"fmt"
	"math"
)

func merhaba(chan1 chan string) {
	chan1 <- "merhaba"
}

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A circle")
}

func area(c circle, mychan chan float64) {
	result := math.Pi * c.r * c.r
	mychan <- result
}

func main() {
	//channel  make() kullanılarak tanımlanır.
	myChannel := make(chan string)
	mychan := make(chan float64)

	go merhaba(myChannel)
	//channel içerisindeki ifadenin yazımı
	fmt.Println(<-myChannel)

	c1 := circle{5}
	c1.display()
	//area fonksiyonundan gelen veriyi channel ile al.
	go area(c1, mychan)
	fmt.Println(<-mychan)
}
