package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type grocery struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var groceries = []grocery{{Name: "Pepsi", Quantity: 1}, {Name: "Coke", Quantity: 4}, {Name: "Sprite", Quantity: 6}}

func getAllGroceries(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(groceries)

}

func getGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["name"]
	for _, val := range groceries {
		if val.Name == data {
			json.NewEncoder(w).Encode(val)
		}
	}

}

func AddGrocery(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	var req grocery
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	groceries = append(groceries, req)
	json.NewEncoder(w).Encode(groceries)
}

func UpdateGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["name"]
	var req grocery
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, val := range groceries {
		if val.Name == data {
			groceries[i].Name = req.Name
			groceries[i].Quantity = req.Quantity
		}
	}
	fmt.Println("req body is", req)
	//groceries = append(groceries, req)
	json.NewEncoder(w).Encode(groceries)
}

func PatchGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["name"]
	var req grocery
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, val := range groceries {
		if val.Name == data {
			if (groceries[i].Name != "") || (groceries[i].Quantity != 0) {
				groceries[i].Name = req.Name
			}
		}
	}
	fmt.Println("req body is", req)
	json.NewEncoder(w).Encode(groceries)
}

func main() {
	router := mux.NewRouter()                                                // for creating http router
	router.HandleFunc("/getAllgroceries", getAllGroceries).Methods("GET")    // for getting all groceries
	router.HandleFunc("/getGrocery/{name}", getGrocery).Methods("GET")       // for getting specific grocery
	router.HandleFunc("/addGrocery", AddGrocery).Methods("POST")             // for adding grocery
	router.HandleFunc("/updateGrocery/{name}", UpdateGrocery).Methods("PUT") // for updating specific grocery
	router.HandleFunc("/patchGrocery/{name}", PatchGrocery).Methods("PATCH") // for updating specific grocery attributes
	fmt.Println("starting http server")
	log.Fatal(http.ListenAndServe(":9191", router)) // for creating http handler which will listen on 9191 port
}
