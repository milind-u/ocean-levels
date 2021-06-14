package coasts

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"

  "github.com/jonas-p/go-shp"

  "util"
)

const (
  url            = "https://maps.googleapis.com/maps/api/elevation/json?path=%v&samples=%v&key=<API_KEY>"
  elevationsFile = "elevations.json"
  // max number of google maps elevation API queries per second
  maxFrequency = 100
  // max size of a google maps elevation API query uri
  maxQuerySize = 8192
)

type location struct {
  Lat float64 `json:"lat"`
  Lng float64 `json:"lng"`
}

type elevation struct {
  Value    float64  `json:"elevation"`
  Location location `json:"location"`
}

func fetchElevations(points []shp.Point) []elevation {
  client := &http.Client{}

  var path strings.Builder
  for i, p := range points {
    // y, x means lat, lng
    path.WriteString(fmt.Sprintf("%v,%v", p.Y, p.X))
    if i != (len(points) - 1) {
      path.WriteString("%7C")
    }
  }
  req, err := http.NewRequest("GET", fmt.Sprintf(url, path.String(), len(points)), nil)
  util.HandleErr(err)

  res, err := client.Do(req)
  util.HandleErr(err)
  defer func() { util.HandleErr(res.Body.Close()) }()

  body, err := ioutil.ReadAll(res.Body)
  var data map[string]json.RawMessage
  util.HandleErr(json.Unmarshal(body, &data))

  var elevations []elevation
  util.HandleErr(json.Unmarshal(data["results"], &elevations))

  return elevations
}

func writeJson(elevations []elevation) {
  body, err := json.MarshalIndent(elevations, "", "  ")
  util.HandleErr(err)
  util.HandleErr(ioutil.WriteFile(elevationsFile, body, 0644))
}
