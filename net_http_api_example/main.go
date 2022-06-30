package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// SOAP API : XML formatında veri dönen api türümüz
// REST API : JSON && XML formatında veri dönen api türümüz

type Car struct {
	ID       int    `xml:"id"`
	Name     string `xml:"name"`
	Price    string `xml:"price"`
	TopSpeed int    `xml:"top_speed"`
}

//method
func (c Car) GetCarName() (name string, price string) {
	return c.Name, c.Price
}

// function
func seyHello() {
	fmt.Println("Selamın aleyküm")
}

type Brand struct {
	ID   int    `json:"id" xml:"ID"`
	Name string `json:"name" xml:"NAME"`
}

func main() {
	// 1 Yol
	/*
		var i320 Car
		i320 = Car{
			ID:       0,
			Name:     "",
			Price:    "",
			TopSpeed: 0,
		}*/
	// 2 Yol
	/*	i320 := Car{
		ID:       1,
		Name:     "BMW",
		Price:    "1282132",
		TopSpeed: 130,
	}*/

	// api setup
	http.HandleFunc("/getCarByQuery", GetCarWhitQuery)
	http.HandleFunc("/getCar", GetCar)
	http.HandleFunc("/getMarka", GetBrand)
	http.HandleFunc("/pullMarka", PullBrand)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func GetCarWhitQuery(w http.ResponseWriter, r *http.Request) {
	var BMW = Brand{
		ID:   1,
		Name: "BMW",
	}
	var BMC = Brand{
		ID:   2,
		Name: "BMC",
	}
	BMWByts, err := json.Marshal(BMW)
	if err != nil {
		return
	}
	BMCByts, err := json.Marshal(BMC)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return
	}
	// multi case
	switch id {
	case 1, 3, 4:
		fmt.Fprintf(w, string(BMWByts))
	case 2:
		fmt.Fprintf(w, string(BMCByts))
	}
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	i320 := Car{
		ID:       1,
		Name:     "BMW",
		Price:    "1282132",
		TopSpeed: 130,
	}
	byts, err := xml.Marshal(i320)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(byts))
}

func GetBrand(w http.ResponseWriter, r *http.Request) {
	marka := Brand{
		ID:   1,
		Name: "Reizler",
	}
	byts, err := xml.Marshal(marka)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(byts))
}

func PullBrand(w http.ResponseWriter, r *http.Request) {
	byts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var marka Brand
	err = json.Unmarshal(byts, &marka)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(marka)
}
