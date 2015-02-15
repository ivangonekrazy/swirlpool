package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/jsonq"
	"os"
)

func xxxmain() {
	// generic data struct
	data := map[string]interface{}{}

	// load JSON string from file
	file, e := os.Open("./json_samples/pullreq.json")
	if e != nil {
		fmt.Printf("Error opening file %v\n", e)
		os.Exit(1)
	}
	defer file.Close()

	// load and decode string into our data struct
	dec := json.NewDecoder(file)
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	fmt.Println(jq.String("action"))
	fmt.Println(jq.String("pull_request", "url"))
	fmt.Println(jq.String("pull_request", "state"))
	fmt.Println(jq.String("pull_request", "user", "login"))
	fmt.Println(jq.String("pull_request", "user", "avatar_url"))
}
