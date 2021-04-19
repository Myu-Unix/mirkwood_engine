package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readConfigPlayer1() error {
	// Read player1.json
	fmt.Println("Reading config file for player1...")
	file, err := os.Open("config/player1.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&MyConfig) // Defined in var
	if err != nil {
		return err
	}
	fmt.Println(MyConfig.race)
	return nil
}
