package gozipcode

import (
    "database/sql"
    _ "fmt"
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
const ASSET_FULLPATH string = "data" + DATABASE_FILENAME

var ZipCodeDBPath string = "/tmp/" + DATABASE_FILENAME
var ZipCodeDB * sql.DB

func Init(){
    // TODO optionally overide ZipCodeDBPath
    // maybe there's a better way to distribute the SQLite file?
    dbdata, err := data.Asset(ASSET_FULLPATH)
    err = ioutil.WriteFile(ZipCodeDBPath, dbdata, 0644)
    ZipCodeDB, err = sql.Open("sqlite3", ZipCodeDBPath)
    if (err != nil){
        return
    }
}

func Cleanup(){
    os.Remove(ZipCodeDBPath)
}

