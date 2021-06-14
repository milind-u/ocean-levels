package coasts

import (
  "log"
  "reflect"
  "time"

  "cln"
  "util"

  "github.com/jonas-p/go-shp"
)

const (
  // total number of shapes in the shp file
  numShapes = 4133
  // max length of a string representation of a latitude or longitude
  maxLatLngLen = 11
  // maximum number of points per google elevation api query
  // Each point is lat, long (max len of 2 * maxLatLngLen), one comma, and one pipe, but no pipe after last point
  maxNumPoints = maxQuerySize/((maxLatLngLen*2)+2) - 1
)

func loadReader() *shp.Reader {
  r, err := shp.Open("coastlines/ne_10m_coastline.shp")
  util.HandleErr(err)
  return r
}

func WriteElevations() {
  r := loadReader()
  defer func() { util.HandleErr(r.Close()) }()
  tq := new(cln.TimeQueue)
  var elevations []elevation
  pct := 0
  for i := float32(0); r.Next(); i++ {
    currentPct := int((i / numShapes) * 100)
    if currentPct > pct {
      pct = currentPct
      log.Printf("%v%% done", pct)
    }
    _, s := r.Shape()

    // convert Shape to PolyLine
    v := reflect.New(reflect.TypeOf(s))
    v.Elem().Set(reflect.ValueOf(s))
    line := *v.Interface().(**shp.PolyLine)

    last := 0
    next := func(j int) int {
      return util.Min(int(line.NumPoints), maxNumPoints+j)
    }
    for j := next(0); j <= int(line.NumPoints) && last < int(line.NumPoints); j = next(j) {
      tq.Add(time.Now())

      elevations = append(elevations, fetchElevations(line.Points[last:j])...)

      // wait until we will not be sending queries to google maps elevation API too fast
      tq.RemoveTimesBefore(time.Now().Add(-time.Second))
      if tq.Len() > maxFrequency {
        time.Sleep(tq.Head().Add(time.Second).Sub(time.Now()))
      }
      _, _ = elevations, last
      last = j
    }
  }
  writeJson(elevations)
  log.Println("done")
}
