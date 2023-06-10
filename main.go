package main

import (
  "fmt"
  "log"
)

func main() {
  var a string

  for {
    fmt.Println("Hurry up!! We are having dinner now! Meatballs...")
    _, err := fmt.Scan(&a)
    if err != nil {
      log.Fatal("Все сломалось")
    }
    if a == "With_pjureshka?" {
      fmt.Println("Yep! With pjureshka!")
      break
    } else {
      fmt.Println("No! With pjureshka! WITH PJURESHKA!")
    }
  }

}
