package main

import (
    "github.com/Hendler/gozipcode"
    "fmt"
)

func main(){
    gozipcode.Init()
    fmt.Print("EXACT MATCH --- \n")
    zipcode := gozipcode.Isequal("04976")
    fmt.Printf("%v\n", *zipcode)
    zipcode = gozipcode.Isequal("adsfa")
    if (zipcode != nil){
        fmt.Printf("%v\n", *zipcode)
    }

    fmt.Print("PREFIX --- \n")
    zipcodes := gozipcode.Islike("0497%")
    for _, zipcode := range zipcodes {
        fmt.Printf("%v\n", *zipcode)
    }

    fmt.Print("RADIUS --- \n")
    skow_lat  := 44.77
    skow_long := -69.71
    zipcodes = gozipcode.Isinradius(skow_lat, skow_long, 15)
    for _, zipcode := range zipcodes {
        fmt.Printf("%v\n", *zipcode)
    }
}