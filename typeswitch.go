package main

import "fmt"

 func main(){
    var x interface{}
    switch i:= x.(type){
       case int:
          fmt.Printf("Integer Type")
       case nil:
          fmt.Printf("Type of x is %T",i)
       case float64:
          fmt.Printf("Float of 64 bits")
       case bool,string:
          fmt.Printf("Bool or String type")
       default:
          fmt.Printf("Type Not Found")
    }
 }
