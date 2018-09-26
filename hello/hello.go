package main

// one-liner like "import ( "fmt" "os" ) doesn't seem to work
import (
  "fmt"
  "os"
  "hello/mymodule"
  "hello/log"
)

type Person struct {
  Name string
  Nationality string
  Single bool
  Steps int
  Direction int
}

// func receiver-decl func-name (args, ...) (return decl, ...) { }
// Note: there's a diff. between p *Person and p Person in receiver-decl (hint: const member)
func (p *Person) Walk(distance int, direction int) (total_steps int, curr_direction int) {
  fmt.Printf("%s walks %d steps toward degree %d\n",
             p.Name, distance, direction)
  p.Steps += distance
  p.Direction = direction
  return p.Steps, p.Direction
}

// inheritance by composition
type Warrior struct {
  Person   // field with type but without a name means inheritance with composition
  Title string
  Weapon string
}

func (w *Warrior) Attack() bool {
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

  guests := [] string { "bobby", "holly", "gates" } // slice  of strings
  // is there a way to do partial initialization?
  guests_detail := [] Person { {"Bobby", "Taiwan", true, 0, 0} } // array of Person objects, without array size declared
  println(guests[0])
  // how to overload println to print a Person object?
  println(guests_detail[0].Name)
  println(len(guests_detail)); // length of arrray

  guests = append(guests, "Elizabeth") // append to slice
  vip_guests := guests[0:2] // create a slice of the array starting at index 0 with size 2

  // iterate over an array
  for index, value := range vip_guests {
    fmt.Printf("%d %s, ", index, value)
  }
  fmt.Printf("\n");

  man := Person { "David", "US", false, 0, 0 }
  total_steps, _ := man.Walk(100, 270) // take only one of the returned values
  total_steps, _ = man.Walk(300, 360)
  fmt.Printf("%s took %d steps\n", man.Name, total_steps)

  // create a composite object
  soldier := Warrior { Person {"Shark", "US", false, 0, 0}, "Cadet", "saber" }
  soldier.Walk(50, 60); // call "inherited" method
  soldier.Attack(); // call its own method

  // swap elements in slice
  stones := [] int {100, 200, 300, 400, 500,}
  stones[2], stones[3] = stones[4], stones[3]

  for index, value := range stones {
    fmt.Printf("[%d] %d \n", index, value)
  }

  mymap := make(map[string]string)
  //this declares nil map which is not initialized and we cannot insert value to it
  var mymap2 map[string] string
  mymap3 :=  map[string] string { "key0": "val0", "key1": "val1", }
  mymap["power"] = "2"
  mymap3["key"] = "hey"
  mymap3["key"] = "heystack"

  fmt.Printf("mymap3[key] is %s\n", mymap3["key"])
  if mymap2 == nil {
    println("mymap2 is nil")
  }

  mymodule.Show(1)

  slogger := log.CreateSimpleLogger()
  slogger.Log("hello log!");

  blogger := log.CreateBeautyLogger()
  blogger.Log("hello log!");

  // use of diff. loggers with same interface
  mymodule.ShowLog(slogger, "show log")
  mymodule.ShowLog(blogger, "show log")

  os.Exit(0)
}

func isOK() (bool, int) {
  return true, 10;
}
