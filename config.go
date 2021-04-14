package main

import (
  "encoding/json"
  "fmt"
  "os"
)

func readConfigPlayer1() {
  // Read player1.json
  fmt.Println("Reading config file for player1...")
  file, _ := os.Open("player1.json")
  defer file.Close()
  decoder := json.NewDecoder(file)
  err := decoder.Decode(&MyConfig) // Defined in var
  if err != nil {
    fmt.Println("error:", err)
  }
  fmt.Println(MyConfig.race)
}