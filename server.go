package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Person struct {
	Name string `json:"nombre"`
	Age  string `json:"edad"`
	Car  string `json:"coche"`
	City string `json:"ciudad"`
}

type People []Person

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.SetCookie(res, &http.Cookie{
			Name:     "my-main-cookie",
			Value:    "Cesar",
			HttpOnly: true,
		})
	})

	//----------------------------
	http.HandleFunc("/vuea", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			//Send string
			res.Write([]byte("GOLANG"))
		}
	})

	//----------------------------
	http.HandleFunc("/vueb", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			//Receive string
			val, _ := ioutil.ReadAll(req.Body)
			fmt.Println(string(val))
		}
	})

	//----------------------------
	http.HandleFunc("/vuec", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			//Sending array of objects
			xx := People{
				Person{Name: "Nuria", Age: "40", Car: "Audi", City: "Santander"},
				Person{Name: "Natalia", Age: "29", Car: "BMW", City: "Santander"},
			}

			json.NewEncoder(res).Encode(xx)
		}
	})

	//-------------------------------
	http.HandleFunc("/vued", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			//Sending object
			json.NewEncoder(res).Encode(&Person{Name: "Natalia", Age: "29", Car: "BMW", City: "Santander"})
		}
	})

	//-------------------------------
	http.HandleFunc("/vuee", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			//Receive object
			var p Person
			json.NewDecoder(req.Body).Decode(&p)
			fmt.Println(p.Name)
			fmt.Println(strconv.Atoi(p.Age))
			fmt.Println(p.Car)
			fmt.Println(p.City)
		}
	})

	//--------------------------------
	http.HandleFunc("/vuef", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			//Receive array of objects
			var p People
			xx, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal(xx, &p)
			for _, i := range p {
				fmt.Println(i.Name + " " + i.Age + " " + i.Car + " " + i.City)
			}
		}
	})

	//------------------------------------
	http.HandleFunc("/vueg", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			//Send array string
			xx := []string{"BMW", "Ferrai", "Audi", "Porsche"}
			res.Write([]byte(strings.Join(xx, ",")))
		}
	})

	//-------------------------------------
	http.HandleFunc("/vueh", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			//Get array string
			val, _ := ioutil.ReadAll(req.Body)
			xx := string(val)
			yy := strings.Split(xx, ",")
			for i := range yy {
				fmt.Println(yy[i])
			}
		}
	})

	http.ListenAndServe(":5000", nil)

}
