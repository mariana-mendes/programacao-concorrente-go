package main

import	("fmt"
  "runtime"
  "time"
  "math")

func main() {

  var m runtime.MemStats
  runtime.ReadMemStats(&m)
  fmt.Println("Alloc = ", bToKb(m.Alloc))
var k float64
k = 1
for k < 14 {
var j float64
j = 0
var x float64
x = math.Pow(2, k)
  for j < x  {
    go func() {
    i := 0
      for i < 10000 {
        i = i + 1
      }
    }()

    go func() {
    i := 0
      for i < 10000 {
        i = i + 1
      }
  }()
  j = j + 1
  }
  t:= time.Duration(5000)  
  time.Sleep(t * time.Millisecond)
  runtime.ReadMemStats(&m)
  fmt.Println("TotalAlloc = ", (bToKb(m.TotalAlloc)))
  fmt.Println("Number of Go routines = ", j)
k = k + 1
}


  runtime.GC()    
}

func bToKb(b uint64) uint64 {
    return b / 1024 
}