# gozipcode


US ZipCode database, batteries included.

go port of https://github.com/buckmaxwell/zipcode


```
 go get github.com/Hendler/gozipcode

```

## usage

Building the example

    go build -o $GOBIN/gozipcode-example example/main.go


```go

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

```

## TODO

- [ ] [zipcodes for all countries](http://stackoverflow.com/questions/308017/where-can-i-get-postal-codes-for-all-countries)
- [ ] tests
- [ ] configurable database location
- [ ] performance - confirm that `check_same_thread` is doing something useful


## building

gozipcode optionally uses [Glide](https://github.com/Masterminds/glide) (`glide.lock`, `glide.yaml` file inculded, `vendor` directory ommitted)

    go get github.com/Masterminds/glide
    go install github.com/Masterminds/glide

    glide install

## dependencies

In `glide.yaml`

    github.com/mattn/go-sqlite3
    github.com/golang/geo

## notes

SQLite file is included in Binary via bindata

    go get -u github.com/jteeuwen/go-bindata
    go install github.com/jteeuwen/go-bindata
    go-bindata -pkg gozipcode -o ./data/zipcode.go ./data/zipcode.db
