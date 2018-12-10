package main


import	("fmt"
         "time"
         "math/rand")
var server1 chan string = make(chan string)
var server2 chan string = make(chan string)
var server3 chan string= make(chan string)

func reliableRequest() string {

  // Sleep para que tenha mais variação de ordem de chegada dos canais.
  rand.Seed(time.Now().UnixNano())
  t:= time.Duration(rand.Intn(200))  
  time.Sleep(t * time.Millisecond)

  select{
    case result := <-server1:
        return "Request by " + result;
    case result := <-server2:
        return "Request by " + result;
    case result := <-server3:
        return "Request by " + result;
  }
}

func main() {
  servers := [3]string{"mirror1.com", "mirror2.br", "mirror3.edu"}

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
