package main

import (
	"testing"
)

func TestRule1CellShouldDiedWhenANeighbourAlive(t *testing.T) {
	var cell life
	cell.alive = true
	cell.rule1(1)
	if cell.alive == LIVE {
		t.Errorf("Expectec cell died but it alive.")
	}
}

func TestRule1CellShouldAliveWhenTwoNeighboursAlive(t *testing.T) {
	var cell life
	cell.alive = true
	cell.rule1(2)
	if cell.alive == DEAD {
		t.Errorf("Expected cell alive but it died.")
	}
}
