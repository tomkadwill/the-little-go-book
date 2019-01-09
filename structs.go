package main

type Saiyan struct {
  Name string
  Power int
}

func main()  {
  goku := Saiyan{
    Name: "Goku",
    Power: 9000,
  }

  println(goku.Name)
  println(goku.Power)

  goku2 := Saiyan{Name: "Boku"}
  goku2.Power = 8000

  println(goku2.Name)
  println(goku2.Power)

  // Using pointers to structs
  gokupointer := &Saiyan{
    Name: "Poku",
    Power: 2000,
  }
  blah(gokupointer)
  println(gokupointer.Power)
}

func blah(s *Saiyan) {
  s.Power += 10000
}
