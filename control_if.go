package main

import "fmt"

func main() {
    var a int = 30
    if a > 20 {
       fmt.Printf("Greater than 20")
    } else if a == 20 {
         fmt.Printf("Equals to 20")
    } else {
    fmt.Printf("Less than 20")
    }
}
