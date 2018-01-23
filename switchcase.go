package main

import "fmt"

func main(){
    var grade string = "B"
    var marks int = 90
    switch marks{
         case 90: grade = "A"
         case 80: grade = "B"
         case 50,60,70: grade = "C"
    }
    switch{
         case grade == "A" :
            fmt.Printf("Excellent")
         case grade == "B" :
            fmt.Printf("Very Good")
         case grade == "C" :
            fmt.Printf("Average")
         default :
            fmt.Printf("Fail")
    }
}
