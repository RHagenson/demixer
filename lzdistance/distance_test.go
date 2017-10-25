package lzdistance

import "testing"

// Ensure the Distance can also be made a float64
func TestDistanceCanBeFloat64(t *testing.T) {
  dist := *new(Distance)
  _ = (float64)(dist)
}
