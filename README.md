# gozipcode


US ZipCode database, batteries included

go port of https://github.com/buckmaxwell/zipcode


```
 go get github.com/Hendler/gozipcode

```

## usage



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