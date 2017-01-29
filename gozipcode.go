package gozipcode

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    _ "log"
    "os"
    "github.com/Hendler/gozipcode/data"
    "io/ioutil"
)

const AUTHOR  string = "Jonathan Hendler"
const LICENSE string = "MIT"
const PACKAGE string = "zipcode"
const VERSION string = "0.9.0"

const DATABASE_FILENAME string = "zipcode.db"
const ASSET_FULLPATH string = "data/" + DATABASE_FILENAME

var ZipCodeDBPath string = "/tmp/" + DATABASE_FILENAME
var ZipCodeDB * sql.DB

func Init(){
    // TODO optionally overide ZipCodeDBPath
    // TODO optionally overwrite file if file exists
    // maybe there's a better way to distribute the SQLite file?
    dbdata, err := data.Asset(ASSET_FULLPATH)
    if (err != nil){
        fmt.Println(err)
        return
    }
    err = ioutil.WriteFile(ZipCodeDBPath, dbdata, 0644)
    if (err != nil){
        fmt.Println(err)
        return
    }
    ZipCodeDB, err = sql.Open("sqlite3", ZipCodeDBPath)
    if (err != nil){
        fmt.Println(err)
        return
    }
}

func Cleanup(){
    os.Remove(ZipCodeDBPath)
}

