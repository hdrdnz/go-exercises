package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID     int
	Kind   string
	Adress string
}

type Interest struct {
	ID   int
	Name string
}

// struct içerisinde struct kullanma örneği
type Person struct {
	ID        int
	Firstname string
	Lastname  string
	Username  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.Firstname + " " + p.Lastname
}

func GetPersonEmailAdress(p *Person, i int) string {
	return p.Email[i].Adress
}

// geriye email nesnesi döner.
func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("****************************")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("fatal error:", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(name string, key interface{}) {
	outfile, err := os.Create(name)
	checkError(err)
	encoder := json.NewEncoder(outfile)
	err = encoder.Encode(key)
	checkError(err)
	outfile.Close()
}

func main() {
	person := Person{
		ID:        8,
		Firstname: "nur",
		Lastname:  "erdem",
		Username:  "nerdem",
		Gender:    "K",
		//ctrl+boşluk yaparsan değerleri görürsün.
		Name: Name{Family: "dijibil", Personal: "nur"},
		Email: []Email{
			Email{ID: 5, Kind: "Work", Adress: "nerdem@spacex.com"},
			Email{ID: 9, Kind: "Home", Adress: "nurerdem@hotmail.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Go"},
			Interest{ID: 2, Name: "Java"},
			Interest{ID: 3, Name: "Ruby"},
		},
	}
	WriteMessage("Reading operation started")
	WriteStarLine()

	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()

	WriteMessage("Personal Email with Index")
	WriteStarLine()
	resemail := GetPersonEmailAdress(&person, 0)
	WriteMessage(resemail)
	WriteStarLine()

	WriteMessage("personel kind ")
	WriteStarLine()
	resemmail2 := GetPersonEmail(&person, 0)
	WriteMessage(resemmail2.Kind)
	WriteStarLine()
	SaveJSON("person.json", person)

}
