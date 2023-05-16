package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	jsonVeri := []byte(`{"İsim":"Ahmet","Soyisim":"Ak","Yas":65}`)
	var goVeri kisi
	//json tipindeki veri go structına dönüştürülür.
	error := json.Unmarshal(jsonVeri, &goVeri)

	if error != nil {
		log.Fatalln(error)
	}

	fmt.Printf("İsim:%s \n Soyisim:%s\n Yas:%d", goVeri.İsim, goVeri.Soyisim, goVeri.Yas)

	if hata != nil {
		panic(hata)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON sonucu:", string(veri))
	fmt.Println(string(veri2))
}
