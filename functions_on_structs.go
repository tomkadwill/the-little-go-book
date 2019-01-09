package main

type Saiyan struct {
  Name string
  Power int
}

func (s *Saiyan) Super() {
  s.Power += 10000
}

func main()  {
  goku := &Saiyan{"Goku", 9001}
  goku.Super()
  println(goku.Power)
}
