# os/exec (Komut Satırına Erişim)

+ os/exec paketi komut satırına (cmd, powershell, terminal) komut göndermemizi sağlayan Golang ile bütünleşik gelen bir pakettir.
+ Bu paket sayesinde oluşturacağımız programa sistem işlerini yaptırabiliriz. Örnek olarak dosya/klasör taşıma/silme/oluşturma/kopyalama gibi işlemleri yaptırabilir. Daha doğrusu komut satırı/terminal üzerinden yapabildiğimiz her işlemi yaptırabiliriz.
+ Komut girerken dikkat etmeniz geren çok önemi bir detay var. Yazacağınız komut birden fazla kelimeden oluşuyorsa mutlaka ayrı ayrı girmelisiniz. Eğer exec.Command() fonksiyonuna direkt olarak “mkdir klasörüm” olarak girseydik, komutu tek kelime olarak algılayacaktı.

+ Eğer Eğer ben hata çıktısının detayını almak istemiyorum, benim işim sadece çıktıyla diyorsanız yapacağımız işlemler :

```go
package main
import (
    "fmt"
    "log"
    "os/exec"
)
func main() {
    cmd := exec.Command("go", "versison")
    çıktı, hata := cmd.CombinedOutput()
    if hata != nil {
        log.Fatalf("Komut hatası: %s\n", hata)
    }
    fmt.Printf(string(çıktı))
}

```