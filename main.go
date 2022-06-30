package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (u User) GetFullName() string {
	return u.Name + " " + u.Surname
}

func main() {
	//var user User // Standart declare
	// user := User{} // short hand
	//--------------------------------------------------//
	/*	user.Name = "faruk"
		user.Surname = "Rakun"
		user.Email = "omer@mail.com"
		user.PhoneNumber = "05555555555"*/
	//--------------------------------------------------//

	http.HandleFunc("/", GetFullName)
	http.HandleFunc("/user", GetUser)
	http.HandleFunc("/users", GetByName)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func GetFullName(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:        "omer",
		Surname:     "tasdemir",
		Email:       "omer@mail.com",
		PhoneNumber: "05555555555",
	}
	fmt.Fprintf(w, user.GetFullName())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:        "omer",
		Surname:     "tasdemir",
		Email:       "omer@mail.com",
		PhoneNumber: "05555555555",
	}
	temp, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(temp))
}

func GetByName(w http.ResponseWriter, r *http.Request) {
	omer := User{
		Name:        "omer",
		Surname:     "tasdemir",
		Email:       "omer@mail.com",
		PhoneNumber: "05555555555",
	}
	sinan := User{
		Name:        "sinan",
		Surname:     "gul",
		Email:       "sinan@mail.com",
		PhoneNumber: "05555555551",
	}
	name := r.URL.Query().Get("name")
	if name == "omer" {
		temp, err := json.Marshal(omer)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, string(temp))
		return
	}
	if name == "sinan" {
		temp, err := json.Marshal(sinan)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, string(temp))
		return
	}

	fmt.Fprintf(w, "Kardeş sen hayırdır")
}
