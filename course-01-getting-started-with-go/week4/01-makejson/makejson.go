package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	//map literal
	//create a map and add the name and address to the map using the keys “name” and “address”, respectively.
	data := map[string]string{
		"name":    "",
		"address": "",
	}

	//prompts the user to enter name and address
	var name string
	var address string

	fmt.Printf("Please enter your name: ")

	nameReader := bufio.NewReader(os.Stdin)
	nameBytes, _, nameErr := nameReader.ReadLine()
	name = string(nameBytes)

	if nameErr == nil {
		data["name"] = name
	}

	fmt.Printf("Please enter your address: ")

	addressReader := bufio.NewReader(os.Stdin)
	addressBytes, _, addressErr := addressReader.ReadLine()
	address = string(addressBytes)

	if addressErr == nil {
		data["address"] = address
	}

	//use Marshal() to create a JSON object from the map, and then print the JSON object.
	barr, err :=
		json.Marshal(data)

	if err == nil {
		fmt.Printf("%s\n", barr)
	}
}
