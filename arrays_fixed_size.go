package main

func main()  {
  // Arrays
  // Fixed in size
  var scores [10]int
  scores[0] = 339

  println(scores[0])

  scores2 := [4]int{9001, 9333, 212, 33}
  for index, value := range scores2 {
    println(index)
    println(value)
  }

  //Slices
  scores3 := []int{1, 4, 293, 4, 9}
  for _, value := range scores3 {
    println(value)
  }

  // 4 ways to initialize a slice
  names := []string{"leto", "jessica", "paul"} // Use when you know the values you want in the slice ahead of time
  checks := make([]bool, 10) // Useful when writing into specific indexes of slices
  var names2 []string // Used in conjunction with append when number of elements is unknown
  scores4 := make([]int, 0, 20) // Useful if you know how many elements you will need

  println(checks)
  println(names2)
  println(scores4)

  // Remember that append can always be using to increase the size of a slice
  newnames := append(names, "tom")
  for _, value := range newnames {
    println(value)
  }
}
