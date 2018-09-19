package main

import "fmt"

type Person struct{
  Name string
  Nationality string
  Single bool
  Steps int
  Direction int
}

func main() {
	fmt.Printf("hello, world\n")
  counter := 1
  fmt.Printf("counter is %d\n", counter)
  ok, val := isOK()
  if (ok) {
    fmt.Printf("OK, val %d\n", val);
  }

  guests := [] string { "bobby", "holly", "gates" }
  guests_detail := [] Person { {"Bobby", "Taiwan", true, 0, 0} }
  println(guests[0])
  println(guests_detail[0].Name);
  println(len(guests_detail));

  man := Person { "David", "US", false, 0, 0 }
  man.Walk(100, 270)
}

func (p *Person) Walk(distance int, direction int) {
  fmt.Printf("%s walks %d steps toward degree %d\n",
             p.Name, distance, direction)
  p.Steps += distance
  p.Direction = direction
}

func isOK() (bool, int) {
  return true, 10;
}
