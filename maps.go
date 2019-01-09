package main

import "fmt"

type Saiyan struct {
  Name string
  Friends map[string]*Saiyan
}

func main()  {
  lookup := make(map[string]int) // This will grow dynamically
  lookup["goku"] = 9001
  power, exists := lookup["vegeta"]
  power2, exists2 := lookup["goku"]
  value := lookup["goku"]

  fmt.Println(power, exists)
  fmt.Println(power2, exists2)
  fmt.Println(value)

  total := len(lookup)
  fmt.Println(total)

  delete(lookup, "goku")
  new_total := len(lookup)
  fmt.Println(new_total)

  lookup2 := make(map[string]int, 100) // Sets default size to 100
  fmt.Println(lookup2)

  goku := &Saiyan{
    Name: "Goku",
    Friends: make(map[string]*Saiyan),
  }

  krillin := &Saiyan{
    Name: "Krillin",
  }

  goku.Friends["Krillin"] = krillin

  lookup3 := map[string]int{
    "goku": 9001,
    "gohan": 2044,
  }

  fmt.Println("loop over keys/values")
  for key, value := range lookup3 {
    fmt.Println(key)
    fmt.Println(value)
  }
}
