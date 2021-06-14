package util

import "log"
import "math"

// HandleErr panics if err is not nil
func HandleErr(err error) {
  if err != nil {
    log.Panicln(err)
  }
}

func intFunc(i, j int, f func(float64, float64) float64) int {
  return int(f(float64(i), float64(j)))
}

// Min returns the smaller value between i and j
func Min(i, j int) int {
  return intFunc(i, j, math.Min)
}

// Max returns the larger value between i and j
func Max(i, j int) int {
  return intFunc(i, j, math.Max)
}
