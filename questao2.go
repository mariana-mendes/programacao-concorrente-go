package main

import	("fmt"

  "time"
  "math/rand")
	var server1 chan string = make(chan string)
	var server2 chan string = make(chan string)
  var server3 chan string= make(chan string)

func reliableRequest() string {


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

  rand.Seed(time.Now().UnixNano())
  started := time.Now()
  t:= time.Duration(rand.Intn(2500))  
  time.Sleep(t * time.Millisecond)

		go func(msg string) {
      server1 <- msg
    }(servers[0])


    go func(msg string) {
      server2 <- msg
    }(servers[1])

    go func(msg string) {
      server3 <- msg
    }(servers[2])
    
    took := time.Since(started)
	  if( (took.Seconds()) > 2){
      fmt.Println("Error")
    }else{
      fmt.Println(reliableRequest());
    }
}
