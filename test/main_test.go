package main

import "testing"

func TestMerhaba(t *testing.T) {
	if Merhaba("Ahmet") != "Merhaba Ahmet" {
		t.Error("Merhaba fonksiyonunda bir sıkıntı var.")
	}
}
