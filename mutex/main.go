package main

import (
	"fmt"
	"sync"
)

// murex, asenkron fonksiyonların içindeki istediğimiz kısımları senkron çalıştırmamızı sağlar.
//mutex,işlemlerin sırayla çalışıtılmasını sağlar.

//mutex için sync paketi kullanılır
//birden fazla asenkron işlem oludğu için waitgroup kullanmayı unutma.
//waitgroup kullanırken pointer yapısını kullan.

var mt sync.Mutex

func paraCek(bakiye *float64, cekilecek float64, wg *sync.WaitGroup) {
	/*
	* mt isimli mutex'i bu işlem yapılırken kilitliyoruz.
	* bu sayede mt mutex'ini başka işlemler kullanamıyor.
	 */
	mt.Lock()

	/*
		bu kısımda asenkron olmasını istemediğimiz işlemi yapalım.
	*/
	*bakiye -= cekilecek
	fmt.Printf("kalan bakiye:%.2f\n", *bakiye)

	/*
	* diğer işlemlerinde kullanabilmesi için mutex'i tekrardan açalım.
	* mt mutex açılınca diğer asenkron işlemdeki mt mutex'i çalışmaya başlar.
	 */
	mt.Unlock()
	fmt.Println("çekme işlemi tamamlandı.")

	//go fonksiyonun bittiğini belirtmek iin yazmalısın.
	wg.Done()

}

func paraYatir(bakiye *float64, yatirilacak float64, wg *sync.WaitGroup) {
	mt.Lock()
	*bakiye += yatirilacak
	fmt.Printf("yeni bakiye:%.2f\n", *bakiye)
	mt.Unlock()
	fmt.Println("yatırma işlemi tamamlandı.")
	wg.Done()
}

func main() {
	/*
	* asenkron işlemlerimizin, ana iş parçacığında tamamlanmasını
	* beklemek için waitgroup nesnesi oluşturalım
	 */
	var wg sync.WaitGroup

	//2 fonksiyonu da bekleyeceğimiz için Add'e 2 yazalım
	wg.Add(2)

	var bakiye float64 = 750
	paraYatir(&bakiye, 50, &wg)
	paraCek(&bakiye, 100, &wg)

	wg.Wait()
	fmt.Println("banka işlemleri tamamlandı.")

}
