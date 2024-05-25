package main

import (
  "bytes"
  "testing"
)

func TestCountWords(t *testing.T) {
  b := bytes.NewBufferString("words1 words2 words3 words4 words5\n")
  expected := 5
  result := Count(b)

  if result != expected {
    t.Errorof("Expected : %d Got : %d\n", expected, reult)
  }
}
