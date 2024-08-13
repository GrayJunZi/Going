package main

import (
	"reflect"
	"testing"
)

func TestEqualPlayers(t *testing.T) {
	expected := Player{
		name: "Mark",
		hp:   100,
	}

	have := Player{
		name: "Alice",
		hp:   59,
	}

	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v but got %+v", expected, have)
	}
}

func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 6
	)

	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("excepted %d but have %d", expected, have)
	}
}
