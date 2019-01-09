package main

import (
  "fmt"
)

type Person struct {
  Name string
}

func (p *Person) Introduce() {
  fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
  *Person
  Power int
}

func main()  {
  person := &Person{
    Name: "Goku",
  }

  goku := &Saiyan{
    Person: person,
    Power: 9001,
  }

  goku.Introduce()
  fmt.Printf(goku.Name)
  fmt.Printf(goku.Person.Name)
}
