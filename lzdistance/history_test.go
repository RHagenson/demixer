package lzdistance

import "testing"

func TestHistoryHasLength(t *testing.T) {
  _ = new(History).Length()
}

func TestHistoryHasContains(t *testing.T) {
  _ = new(History).Contains()
}

func TestHistoryHasAdd(t *testing.T) {
  _ = new(History).Add(1)
}
