package main

import (
	"encoding/json"
	"fmt"
)

type kisi struct {
	İsim    string
	Soyisim string
	Yas     int64
}

func main() {
	murat := kisi{
		İsim:    "Murat",
		Soyisim: "Erdem",
		Yas:     25,
	}
	//go tipini json tipine dönüştürülür.
	veri, hata := json.Marshal(murat)
	veri2, err := json.MarshalIndent(murat, " 1", "    ")

	if hata != nil {
		panic(hata)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON sonucu:", string(veri))
	fmt.Println(string(veri2))
}
