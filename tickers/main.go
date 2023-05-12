package main

//zamanlayıcılar, belirli sürede bir tekrar etme işlemi için kullanılır.
//Zamanlayıcılar programın çalışma süresince veya durdurulana kadar çalışabilir.
// Select ile çoklu goroutine işlemlerinin iletişimini bekleyebiliriz
import (
	"fmt"
	"time"
)

func main() {
	//yarım saniyede bir mesaj tetikler
	//mesj için newTicker ifadesi kullanılır.
	tekrar := time.NewTicker(500 * time.Millisecond)
	bitti := make(chan bool)
	//anonim goroutine fonksiyonu oluşturduk.
	go func() {
		//sonsuz for döngüsü
		for {
			//select ile kanal iletişimleri dinlenilir.
			select {
			//bitti ifadesi true olursa döngü biter.
			case <-bitti:
				return
				//tekrar.C ile zaman bilgisi alınır.
			case zaman := <-tekrar.C:
				fmt.Println("tekrar zamanı:", zaman)
			}
		}
	}()
	//virgüllü saniye işlemleri için millisecond ifadesini kullan.
	time.Sleep(1600 * time.Millisecond) //1,6 saniye programı uyuttuk.
	//time.Sleep(1 * time.Second)
	tekrar.Stop() //tekrarlama işlemini durdurduk.
	bitti <- true
	fmt.Println("tekrarlayıcı durdu")
}
