package main

import (
    "github.com/Hendler/gozipcode"
    "fmt"
)

func main(){
    gozipcode.Init()
    zipcode := gozipcode.Isequal("04976")
    fmt.Printf("%v\n", *zipcode)
    zipcodes := gozipcode.Islike("0497%")
    for _, zipcode := range zipcodes {
        fmt.Printf("%v\n", *zipcode)
    }
}