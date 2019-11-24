package main

import (
  "fmt"
  "github.com/hpaasch/stringutil"
)

func main() {
  cows := "Big cows hide here."
  fmt.Println("cows:", stringutil.Reverse(cows))

  pigs := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
  fmt.Println("pigs:", pigs)

  for i:=0; i < len(pigs); i++ {
    fmt.Printf("%x ", pigs[i])
  }

  fmt.Println("%x\n", pigs)

  testSlice := "dog is good"
  for i:=0; i<len(testSlice); i++ {
    fmt.Println(testSlice[i])
  }

  for _, letter := range testSlice {
    fmt.Println("it:", letter)
  }

  sum := 1
  for sum < 5 {
    fmt.Println(sum)
    sum += 1
  }

  slice := make([]int, 10)
  for i, _ := range slice {
    slice[i] = 1 << uint(i)
  }
  fmt.Println(slice)

}
