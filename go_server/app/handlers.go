package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerList wraps the array of customers for proper XML encoding
type CustomerList struct {
	XMLName   xml.Name   `xml:"customers"`
	Customers []Customer `xml:"customer"`
}

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city"	xml:"City"`
	Zipcode string `json:"Zipcode"	xml:"Zip_code"`
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customer := CustomerList{
		Customers: []Customer{
			{Name: "Prathap", City: "Pune", Zipcode: "411013"},
			{Name: "Raj", City: "Delhi", Zipcode: "411000"},
		},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		w.Write([]byte(xml.Header))
		fmt.Print("xml")
		if err := xml.NewEncoder(w).Encode(customer); err != nil {
			// Handle error case
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Print("json")
		json.NewEncoder(w).Encode(customer)
	}

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!!")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
