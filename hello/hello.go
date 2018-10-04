package main

// one-liner like "import ( "fmt" "os" ) doesn't seem to work
import (
  "fmt"
  "os"
  "unsafe"
  "hello/mymodule"
  "hello/log"
  "time"
  "sync"
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
  mymodule.Show(2)

  slogger := log.CreateSimpleLogger()
  slogger.Log("hello log!");

  blogger := log.CreateBeautyLogger()
  blogger.Log("hello log!");

  // use of diff. loggers with same interface
  mymodule.ShowLog(slogger, "show log")
  mymodule.ShowLog(blogger, "show log")

  demo_defer(2)

  // initialized if: if init-expression; condition { }
  if _, r:=isOK(); r>3 {
    fmt.Printf("inited if is OK!\n")
  }

  // declare multiple variables in if-init
  if a, b, c := 3, true, 4; a==3 && b {
    fmt.Printf("inited if 2 is OK! %d\n", c);
  }

  var a = Person {"Shark", "US", false, 0, 0}
  a.Name = "Bison"
  // why this doesn't work and need newPerson to return an object?
  //if a := Person {"Shark", "US", false, 0, 0} ; a.Name == "Shark" {
  if a:= newPerson("Shark"); a.Name == "Shark" {
    fmt.Printf("inited if 3 is %s!\n", a.Name)
  }

  // strings and byte arrays
  str := "word"
  ba := [] byte (str)
  fmt.Printf("size of byte array from %s(size %d) is %d\n", str, len(str), len(ba)) // finally no null-terminated crap

  // string and rune
  for _, value := range str {
    fmt.Printf("size of str[0], which is a rune: %d\n", unsafe.Sizeof(value))
    break
  }

  // function type
  lg := LogPrepend { BLog }
  lg.logger("hello")

  lg2 := LogPrepend { YLog }
  lg2.logger("hello")
  lg2.logger = BLog
  lg2.logger("hello")

  // how to wait for go routines to finish
  var wg sync.WaitGroup
  wg.Add(2)

  // go routines with anonymous routine
  go func () {
    defer func() {
      fmt.Printf("go routine done\n")
      wg.Done()
    } ()
    fmt.Printf("start going!\n")
    for i := 0; i<10; i++ {
      time.Sleep(100*time.Millisecond)
      fmt.Printf("go %d!\n", i)
    }
  } ()

  // create a channel
  c := make(chan int)

  // note that WaitGroup must not be passed by value or
  // it gets copied into another object and cause fatal error on wg.Wait()
  go my_routine(c , &wg)

  // feed data to go routines via channel
  for i := 5; i > -2; {
    select {
    case  c <- i:
      fmt.Printf("sending %d\n", i)
      i--
    default:
    }
  }

  fmt.Printf("waiting for go routines to finish\n")
  wg.Wait()
  fmt.Printf("all go routines finished\n")
  // no need for waiting
  // time.Sleep(2*time.Second)

  os.Exit(0)
}

func my_routine (c chan int, wg *sync.WaitGroup) {
  defer wg.Done()
  for {
    r := <- c
    if r<0 {
      fmt.Printf("myroutine break\n")
      break
    } else {
      fmt.Printf("my_routine %d\n", r)
    }
  }
  fmt.Printf("myroutine exits\n")
}

// function type
type Log func (message string)

type LogPrepend struct {
  logger Log
}

func BLog (message string) {
  fmt.Printf("[basic] %s\n", message)
}

func YLog (message string) {
  fmt.Printf("[yarn] %s\n", message)
}

func newPerson(name string) Person {
  return Person {name, "US", false, 0, 0}
}

func isOK() (bool, int) {
  return true, 10;
}

func demo_defer(i int) () {
  // defer is like a stack unwind on exception or scope guard. So convenient!
  defer fmt.Printf("defer1\n")
  defer func () {
    fmt.Printf("defer2\n")
  } () // note the trailing ()

  // cannot defer any expression, but enclose the expression in a function call
  defer func () { if i<3 { fmt.Printf("defer3\n") } } ()
}
