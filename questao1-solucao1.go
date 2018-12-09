package main

import "fmt"

var canal chan string = make(chan string, 1)

func request(serverName string){
     canal <- serverName
}

func reliableRequest(){

    go func(msg string) {
        request(msg)
      }("mirror1.com")

      go func(msg string) {
        request(msg)
      }("mirror2.br")

      go func(msg string) {
        request(msg)
      }("mirror3.edu")

  fmt.Println("chegou primeiro: ", <- canal)
}


func main() {
 reliableRequest()
 
}

