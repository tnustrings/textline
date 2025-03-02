package main

import "fmt"


func Hello() string {
  message := fmt.Sprintf("hi")
  return message
}


// main runs Hello.
func main() {
  fmt.Println(Hello())
}



