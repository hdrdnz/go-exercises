# Ekstra Notlar

## Dışa Aktarma (Exporting)
+ Daha önce diğer programlama dillerinde public anahtar kelimesi olarak gördüğümüz dışa aktarma işlemi golang'ta ise,dışa aktarılmasını istediğimiz öğeyi oluştururken baş harfini büyük yazarız

## Çok Satırlı String Oluşturma
+ Çok satırlı string oluşturmak için (`) işaretini kullanırız.
+ Türkçe klavyeden alt gr ve virgül tuşuna basarak bu işaret kullanılır.

## Çapraz Platform Dosya Yolları
+ Çapraz Platform: birden fazla işletim sisteminde dağıtımı olan yazılımlardır.
+ Bir işletim sisteminde dosyanın veya dizinin yolunu belirtmek için taksim veya ters-taksim işaretleri kullanırız. Fakat yazağımız program çapraz-platformsa ise <b>os.PathSeperator </b>kullanabiliriz.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	s := string(os.PathSeparator)
	yol := "dosyalar" + s + "muzikler"
	fmt.Println(yol)
}

```

## İşletim Sistemini Görme

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {

	if r := runtime.GOOS; r == "windows" {
		fmt.Println("Windows için yönetici olarak çalıştırın.")
	} else if r == "linux" {
		fmt.Println("Linux için sudo komutu ile çalıştırın.")
	} else {
		fmt.Println("Geçersiz işletim sistemi!")
	}
}
```