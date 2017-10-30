package ami

import "testing"

// Ensure the Distance can also be made a float64
func TestProbabilityCanBeFloat64(t *testing.T) {
  dist := *new(Probability)
  _ = (float64)(dist)
}
