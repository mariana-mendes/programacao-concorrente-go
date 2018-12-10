
package main

import	("fmt"
  "runtime"
  "time"
  "math/rand")

func main() {

  var m runtime.MemStats
  runtime.ReadMemStats(&m)
  fmt.Println("Alloc = ", bToKb(m.Alloc))
  j := 0
  for j < 10 {
    fmt.Println("Alloc = ", bToKb(m.Alloc))
    runtime.ReadMemStats(&m)
    go func() {

      i := 0
      for i < 10 {
        t:= time.Duration(rand.Intn(100))  
        time.Sleep(t * time.Millisecond)
        i = i + 1
      }
    }()

    go func() {
    i := 0
    for i < 10 {
      
      runtime.ReadMemStats(&m)
      t:= time.Duration(rand.Intn(100))  
      time.Sleep(t * time.Millisecond)
      i = i + 1
    }
  }()


  j = j + 1
  }
  t:= time.Duration(5000)  
  time.Sleep(t * time.Millisecond)
  fmt.Println("TotalAlloc = ", bToKb(m.TotalAlloc))

  runtime.GC()    
}

func bToKb(b uint64) uint64 {
    return b / 1024 
}