package gozipcode

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    _ "log"
    "os"
    "github.com/Hendler/gozipcode/data"
    "io/ioutil"
    "math"
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
    // TODO optionally overide ZipCodeDBPath (can already by done since it's public)
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

    // we don't want threadsafe
    // not sure that `check_same_thread=false` is respected as with python sqlite
    ZipCodeDB, err = sql.Open("sqlite3", ZipCodeDBPath + "?check_same_thread=false")
    if (err != nil){
        fmt.Println(err)
        return
    }
}

func Cleanup(){
    os.Remove(ZipCodeDBPath)
}


const kmtomiles = float64(0.621371192)
const earthRadius = float64(6371)

/*
 * The haversine formula will calculate the spherical distance as the crow flies
 * between lat and lon for two given points in km
 * from https://play.golang.org/p/MZVh5bRWqN

    func main() {
        var locationName [2]string
        var location [2][2]float64
        // York - lat,lon
        locationName[0] = "York"
        location[0][0] = 1.0803
        location[0][1] = 53.9583
        // Bristol - lat,lon
        locationName[1] = "Bristol"
        location[1][0] = 2.5833
        location[1][1] = 51.4500

        // Use haversine to get the resulting diatance between the two values
        var distance = Haversine(location[0][0], location[0][1], location[1][0], location[1][1])
        // We wish to use miles so will alter the resulting distance
        var distancemiles = distance * kmtomiles

        fmt.Printf("The distance between %s and %s is %.02f miles as the crow flies", locationName[0], locationName[1], distancemiles)
    }

 */
func Haversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
    var deltaLat = (latTo - latFrom) * (math.Pi / 180)
    var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)
    var a = math.Sin(deltaLat / 2) * math.Sin(deltaLat / 2) +
        math.Cos(latFrom * (math.Pi / 180)) * math.Cos(latTo * (math.Pi / 180)) *
        math.Sin(deltaLon / 2) * math.Sin(deltaLon / 2)
    var c = 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
    distance = earthRadius * c

    return
}

//
type Zipcode struct {
    // The 5 digit zip code
    ZIP_CODE string
    // The type of zip code according to USPS: 'UNIQUE', 'PO BOX', or 'STANDARD'"""
    ZIP_CODE_TYPE string
    // The primary city associated with the zip code according to USPS
    CITY string
    // The state associated with the zip code according to USPS
    STATE string
    // This value will always be 'Primary'. Secondary and 'Not Acceptable' placenames have been removed.
    LOCATION_TYPE string
    // The latitude associated with the zipcode according to the National Weather Service.  This can be empty when there is no NWS Data
    LAT float64
    // The longitude associated with the zipcode according to the National Weather Service. This can be empty when there is no NWS Data
    LONG float64
    XAXIS string
    YAXIS string
    ZAXIS string
    // This value will always be NA for North America
    WORLD_REGION string
    // This value will always be US for United States -- This includes Embassy's, Military Bases, and Territories
    COUNTRY string
    // The city with its state or territory. Example:  'Cleveland, OH' or 'Anasco, PR'
    LOCATION_TEXT string
    // "A string formatted as WORLD_REGION-COUNTRY-STATE-CITY. Example: 'NA-US-PR-ANASCO'
    LOCATION string
    // A boolean value that reveals if a zipcode is still in use
    DECOMMISIONED string
    // Number of tax returns filed for the zip code in 2008 according to the IRS
    TAX_RETURNS_FILED string
    // Estimated population in 2008 according to the IRS
    ESTIMATED_POPULATION string
    // Total wages according in 2008 according to the IRS"
    TOTAL_WAGES string
    // Not empty when there is no NWS data.
    NOTES string
}

// Takes a partial zip code and returns a list of zipcode objects with matching prefixes.
func Islike(zipcode string) []*Zipcode {
    valid := validate(zipcode)
    if (!valid){
        return nil
    }
    rows, err := ZipCodeDB.Query("SELECT * FROM ZIPS WHERE ZIP_CODE LIKE ?", zipcode)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer rows.Close()
    zipcodes := make([]*Zipcode, 0)
    for rows.Next() {
        zip_code := new(Zipcode)
        err = rows.Scan(&zip_code.ZIP_CODE,
            &zip_code.ZIP_CODE_TYPE,
            &zip_code.CITY,
            &zip_code.STATE,
            &zip_code.LOCATION_TYPE,
            &zip_code.LAT,
            &zip_code.LONG,
            &zip_code.XAXIS,
            &zip_code.YAXIS,
            &zip_code.ZAXIS,
            &zip_code.WORLD_REGION,
            &zip_code.COUNTRY,
            &zip_code.LOCATION_TEXT,
            &zip_code.LOCATION,
            &zip_code.DECOMMISIONED,
            &zip_code.TAX_RETURNS_FILED,
            &zip_code.ESTIMATED_POPULATION,
            &zip_code.TOTAL_WAGES,
            &zip_code.NOTES)
        if err != nil {
            fmt.Println(err)
            return nil
        }
        zipcodes = append(zipcodes, zip_code)
    }
    return zipcodes
}


func Isequal(zipcode string) *Zipcode{
    valid := validate(zipcode)
    if (!valid){
        return nil
    }
    row := ZipCodeDB.QueryRow("SELECT * FROM ZIPS WHERE ZIP_CODE == ?", zipcode)
    zip_code := new(Zipcode)
    err := row.Scan(&zip_code.ZIP_CODE,
                    &zip_code.ZIP_CODE_TYPE,
                    &zip_code.CITY,
                    &zip_code.STATE,
                    &zip_code.LOCATION_TYPE,
                    &zip_code.LAT,
                    &zip_code.LONG,
                    &zip_code.XAXIS,
                    &zip_code.YAXIS,
                    &zip_code.ZAXIS,
                    &zip_code.WORLD_REGION,
                    &zip_code.COUNTRY,
                    &zip_code.LOCATION_TEXT,
                    &zip_code.LOCATION,
                    &zip_code.DECOMMISIONED,
                    &zip_code.TAX_RETURNS_FILED,
                    &zip_code.ESTIMATED_POPULATION,
                    &zip_code.TOTAL_WAGES,
                    &zip_code.NOTES)
    if err == sql.ErrNoRows {
        fmt.Println(err)
        return nil
      } else if err != nil {
        fmt.Println(err)
        return nil
      }
    return zip_code

}



func validate(zipcode string) bool {
    for _, value := range zipcode {
        switch {
        case value >= '0' && value <= '9':
            return true
        }
    }
    return false
}


// def _validate(zipcode):
//     if not isinstance(zipcode, str):
//         raise TypeError('zipcode should be a string')
//     int(zipcode) # This could throw an error if zip is not made of numbers
//     return True



// def isinradius(point, distance):
//     """Takes a tuple of (lat, lon) where lon and lat are floats, and a distance in miles. Returns a list of zipcodes near the point."""
//     zips_in_radius = list()

//     if not isinstance(point, tuple):
//         raise TypeError('point should be a tuple of floats')
//     for f in point:
//         if not isinstance(f, float):
//             raise TypeError('lat and lon must be of type float')

//     dist_btwn_lat_deg = 69.172
//     dist_btwn_lon_deg = math.cos(point[0]) * 69.172
//     lat_degr_rad = float(distance)/dist_btwn_lat_deg
//     lon_degr_rad = float(distance)/dist_btwn_lon_deg

//     latmin = point[0] - lat_degr_rad
//     latmax = point[0] + lat_degr_rad
//     lonmin = point[1] - lon_degr_rad
//     lonmax = point[1] + lon_degr_rad

//     if latmin > latmax:
//         latmin, latmax = latmax, latmin
//     if lonmin > lonmax:
//         lonmin, lonmax = lonmax, lonmin

//     stmt = ('SELECT * FROM ZIPS WHERE LONG > {lonmin} AND LONG < {lonmax}\
//      AND LAT > {latmin} AND LAT < {latmax}')
//     _cur.execute(stmt.format(lonmin=lonmin, lonmax=lonmax, latmin=latmin, latmax=latmax))
//     results = _cur.fetchall()

//     for row in results:
//         if haversine(point, (row[_LAT], row[_LONG])) <= distance:
//             zips_in_radius.append(Zip(row))
//     return zips_in_radius



















