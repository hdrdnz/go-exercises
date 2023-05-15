package main

import (
	"log"
	"os"
)

// ilk olarak log paketimiz ile ilgili ön ayarlamalar yapılır.
func init() {
	//SetPrefix() fonksiyonu ile log çıktımızın satırının başında ne yazacağını belirleyebiliyoruz.
	log.SetPrefix("kayıt:")
	//log.Ldate bize zamanını gösteriyor.
	//log.Llongfile ise dosyanın bulunduğu dizin ile ismini ve yapılan işlem ile ilgili satırı gösteriyor.
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

}
func main() {
	//log.Println() fonksiyonu ile klasik log çıktılama işlemi
	log.Println("main fonksiyonu başladı.")

	//log.Fatalln() fonksiyonu ile kritik hataları bildirir. log.Println() fonksiyonundan farkı program 1 çıkış kodu ile biter. Bu da programın hatalı bittiği anlamına gelir.
	log.Fatalln("ölümcül hata")

	//log.Panicln() fonksiyonunda ise ekrana çıktımızı verir ve aynı zamanda bunu normal panic() fonksiyonu ile yapar.
	//panic mesajı programı direkt sonlandırır.
	//log.Panicln("PANİK")

	log.SetPrefix("Kayıtım:")
	// log.Lshortfile yazarak dosyamızın ismini dizin olmadan ve çıktının yapıldığı satırı bastırmasını istedik.
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// loglarımız komut satırında gözükmek yerine belirttiğimiz dosyaya yazılacaktır.
	logDosyam, _ := os.Create("kayıt.log")
	defer logDosyam.Close()
	log.SetOutput(logDosyam)
	log.Println("bir sıkıntı yaşadık ama mühim değil.")
	log.Fatal("bu sefer kötü bir şey oldu.")

}
