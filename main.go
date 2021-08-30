package main

import (
    "encoding/json"
    "fmt"
    "os"
    "io/ioutil"
)

type Gpx struct {
  Points []string `json:"points"`
}

func main() {
  // Open our jsonFile
  jsonFile, err := os.Open("/mnt/c/Users/me/Downloads/mapstogpx20210829134635.gpx")
  // if we os.Open returns an error then handle it
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Successfully Opened file")
  // defer the closing of our jsonFile so that we can parse it later on
  defer jsonFile.Close()

  // read our opened jsonFile as a byte array.
  byteValue, _ := ioutil.ReadAll(jsonFile)

  var gpx Gpx
  // we unmarshal our byteArray which contains our
  // jsonFile's content into 'users' which we defined above
  json.Unmarshal(byteValue, &gpx)

  fmt.Println("latitude,longitude,name")

  for i, pts := range(gpx.Points) {
    fmt.Printf("%s,point %d\n", pts, i)
  }

}
