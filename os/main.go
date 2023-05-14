package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	//exec.Command() fonksiyonuyla komutlarımızı girdik.
	//komutu tanımladık.
	cmd := exec.Command("go", "version")

	//cmd.StdoutPipe() fonksiyonuyla gönderdiğimiz komutun çıktılarını alabiliyoruz
	cmdOkuyucu, hata := cmd.StdoutPipe()
	//cmdOkuyucu değişkenine komut çıktımızı aldık

	if hata != nil {
		fmt.Fprintln(os.Stderr, "ÇIKTI OKUNURKEN HATA OLUŞTU", hata)
		//1 numaralı çıkış kodu hatalar için kullanılır
		os.Exit(1)
	}

	// çıktıyı okumak için çıktı komutu tanımlandı.Bu değişkenimiz cmdOkuyucu değişkenini taramaya yarayacak.
	cıktı := bufio.NewScanner(cmdOkuyucu)
	go func() {
		for cıktı.Scan() {
			fmt.Println(cıktı.Text())
		}
	}()

	hata = cmd.Start()

	if hata != nil {
		fmt.Fprintln(os.Stderr, "hata çalışırken hata oluştu", hata)
		os.Exit(1)
	}

	hata = cmd.Wait()

	if hata != nil {
		fmt.Fprintln(os.Stderr, "beklerken hata olutşu", hata)
		os.Exit(1)
	}

}
