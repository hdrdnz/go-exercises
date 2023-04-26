package main

//eş zamanlı=Concurrency
// goroutine :eş zamanlı olarak yapılan görevlerin her biridir.
//goroutine eş zamanlı olarak çalışır.
//goroutine birden fazla işin aynı anda yapılmasını sağlar.
//go anahtar kelimesi kullanılır.Bir fonksiyonun başına go terimini eklersek artık bu fonksiyona yeni bir thread üzerinde asenkron bir şekilde arka planda çalış demiş oluruz.
//main go rutini main fonk.çalıştığı an çalışmaya başlar.
// !!goroutine arasındaki geçişler kontrol edilemez.Hangisinin önce çalışacağını belirleyemezsin.
//main fonksiyonu içerisinde bütün gorotuine'lerin verilen bekleme süresinde tamamlanıp tamamlanmayacağı bilinemez.
import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup içerisinde add, wait,done metotları bulunur.
var wg sync.WaitGroup

func main() {
	// main fonksiyonu çalışmadan önce 1 tane goroutine çalışması gerekiyor.
	wg.Add(2)
	go printX()
	go printY()
	wg.Wait()

	//fmt.Println()

	//main fonksiyonunun bitmesini 1 saniye geciktirir. fakat bu yöntemi kullanmak efektif değildir çünkü go hızlı bir dildir.1000'lerce isteğe hizmet verecek üretim sınıfı bir sunucu yaparsanız, bu  şekilde bloke etmek büyük bir sorun olacaktır.
	time.Sleep(time.Second)
}

func printX() {
	for i := 0; i < 500; i++ {
		fmt.Print("X")
	}
	wg.Done()
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
	wg.Done()
}
