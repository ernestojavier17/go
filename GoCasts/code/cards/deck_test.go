package main

import "testing"

func TestNewDeck_lengthOfDeckIs4(t *testing.T) {

	d := newDeck()

	if len(d) != 17 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))

	}



}
