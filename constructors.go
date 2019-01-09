package main

type Saiyan struct {
  Name string
  Power int
  Father *Saiyan
}

func NewSaiyan(name string, power int) *Saiyan {
  return &Saiyan{
    Name: name,
    Power: power,
  }
}

// func NewSaiyan(name string, power int) Saiyan {
//   return Saiyan{
//     Name: name,
//     Power: power,
//   }
// }

func main()  {
  goku := new(Saiyan)
  goku.Name = "Tom"
  goku.Father = NewSaiyan("Dad", 10)
  println(goku.Name)
  println(goku.Father.Name)

  goku2 := NewSaiyan("Dave", 2)
  println(goku2.Name)
}
