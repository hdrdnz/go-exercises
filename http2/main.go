// api uygulaması
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// çıktı olarak json tipinde çıktı vereceğiz.
type API struct {
	Message string "json:message"
}

type User struct {
	Id        int    "json:id"
	Firstname string "json:firstname"
	Lastname  string "json:lastname"
	Age       int    "json:age"
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error:", err.Error())
		os.Exit(1)
	}
}

func main() {
	//url de ilk bu kısım yazacak.
	apiRoot := "/api"

	//.../api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API home"}
		//go structı json tipine dönüştürülür.
		output, err := json.Marshal(message)
		checkError(err)
		//json tipi byte olduğu için stringe çevrilmesi gerekiyor.
		fmt.Fprintf(w, string(output))
	})

	//../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{Id: 1, Firstname: "Havva Nur", Lastname: "Durudeniz", Age: 22},
			User{Id: 2, Firstname: "Murat Can", Lastname: "Tanrıverdi", Age: 25},
		}
		message := users
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	//../api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{Id: 1, Firstname: "Havva Nur", Lastname: "Durudeniz", Age: 22}
		output, err := json.Marshal(user)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	//api uygulamasını çalıştırmak için:
	http.ListenAndServe(":8080", nil)

}
