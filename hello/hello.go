package main

// one-liner doesn't seem to work
import (
  "fmt"
  "os"
)

type Person struct {
  Name string
  Nationality string
  Single bool
  Steps int
  Direction int
}

// func receiver-decl func-name (args, ...) (return decl, ...) { }
// Note: there's a diff. between p *Person and p Person in receiver-decl
func (p *Person) Walk(distance int, direction int) (total_steps int, curr_direction int) {
  fmt.Printf("%s walks %d steps toward degree %d\n",
             p.Name, distance, direction)
  p.Steps += distance
  p.Direction = direction
  return p.Steps, p.Direction
}

// inheritance by composition
type Warrior struct {
  Person   // field with type but without a name
  Title string
  Weapon string
}

func (w *Warrior) Attack() (bool) {
  fmt.Printf("%s %s attacks with %s\n", w.Title, w.Name, w.Weapon)
  return true; // attack is successful
}

func main() {

  if len(os.Args) < 2 {
    fmt.Printf("Usage: %s arg1\n", os.Args[0])
  }

	fmt.Printf("hello, world\n")
  counter := 1 // declare an variable and assign value to it
  var counter2 int  // declare a variable, which takes zero initial value
  fmt.Printf("counter is %d, 2nd is %d\n", counter, counter2)
  counter = 3; // assign value to an object
  ok, val := isOK() // multiple return values
  if (ok) {
    fmt.Printf("OK, val %d\n", val);
  }

  guests := [] string { "bobby", "holly", "gates" } // simple array of strings
  // is there a way to do partial initialization?
  guests_detail := [] Person { {"Bobby", "Taiwan", true, 0, 0} } // array of Person objects
  println(guests[0])
  // how to overload println to print a Person object?
  println(guests_detail[0].Name);
  println(len(guests_detail)); // length of arrray

  man := Person { "David", "US", false, 0, 0 }
  total_steps, _ := man.Walk(100, 270) // take only one of the returned values
  total_steps, _ = man.Walk(300, 360)
  fmt.Printf("%s took %d steps\n", man.Name, total_steps)

  // create a composite object
  soldier := Warrior { Person {"Shark", "US", false, 0, 0}, "Cadet", "saber" }
  soldier.Walk(50, 60); // call "inherited" method
  soldier.Attack(); // call its own method

  os.Exit(0)
}

func isOK() (bool, int) {
  return true, 10;
}
