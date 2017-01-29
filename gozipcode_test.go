package gozipcode

import (
    "testing"
    "os"
)


func setup(){
    Init()
}

func tearDown(){

}

func TestMain(m *testing.M) {
    setup()
    retCode := m.Run()
    tearDown()
    os.Exit(retCode)
}

func TestNilResults(t *testing.T){
    bad_string := "049734534534%"
    zipcodes :=  Islike(bad_string)

    if len(zipcodes) > 0  {
        t.Error("zipcodes should be nil or empty Islike")
    }

}

func TestNilResult(t *testing.T){
    bad_string := "adsfa"
    zipcode :=  Isequal(bad_string)

    if zipcode != nil {
        t.Error("zipcode should be nil Isequal")
    }
}

func TestSingleResult(t *testing.T){
    good_string := "04976"
    zipcode :=  Isequal(good_string)

    if zipcode.ZIP_CODE != good_string {
        t.Error("zipcode should be 04976")
    }
}


