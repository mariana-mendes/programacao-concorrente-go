package main

import	"fmt"
	var server1 chan string = make(chan string)
	var server2 chan string = make(chan string)
  var server3 chan string= make(chan string)

func reliableRequest() string {
  select{
    case result := <-server1:
        return "Request by" + result;
    case result := <-server2:
        return "Request by" + result;
    case result := <-server3:
        return "Request by" + result;
  }

}

func main() {
  servers := [3]string{"mirror1.com", "mirror2.br", "mirror3.edu"}

  // o for eh so pra ficar melhor de visualizar que a ordem de chegada das requisições pode mudar
  for i := 0; i < 10; i++ {
		  go func(msg string) {
      server1 <- msg
    }(servers[0])


    go func(msg string) {
      server2 <- msg
    }(servers[1])

    go func(msg string) {
      server3 <- msg
    }(servers[2])

    fmt.Println(reliableRequest());
  }
}
