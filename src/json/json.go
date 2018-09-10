package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Country string `json:"country"`
	DOB     string `json:"dob"`
}

func main() {

	p := &Person{
		Name:    "phani",
		Id:      "CP11786",
		Country: "India",
		DOB:     "AUG-09-1988",
	}

	bytes, err := json.Marshal(p)

	if err != nil {
		fmt.Printf("Error while marshalling the json %s\n", err)
	}

	fmt.Println(string(bytes))

	var p1 Person

	err1 := json.Unmarshal(bytes, &p1)
	if err != nil {
		fmt.Printf("Error occured while unmarshalling %s\n", err1)
	}

	fmt.Println(p1.Country)
	fmt.Println(p1.Name)
	fmt.Println(p1.Id)

}
