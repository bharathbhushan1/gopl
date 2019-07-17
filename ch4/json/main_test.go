package main

import (
	"bytes"
	"testing"
)

func TestMarshal(t *testing.T) {
	want := []byte(`[{"Title":"a","released":2000,"color":true,"Actors":["x1","x2","x3"]},{"Title":"b","released":2001,"Actors":["y1","y2","y3"]}]`)
	if got := Marshal(movies); !bytes.Equal(got, want) {
		t.Errorf("%q %q", got, want)
	}
}
