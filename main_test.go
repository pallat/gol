package main

import (
	"testing"
)

func TestRule1CellShouldDiedWhenANeighbourAlive(t *testing.T) {
	var cell life
	cell.alive = LIVE
	cell.rule(1)
	if cell.alive == LIVE {
		t.Errorf("Expectec cell died but it alive.")
	}
}

func TestRule1CellShouldAliveWhenTwoNeighboursAlive(t *testing.T) {
	var cell life
	cell.alive = LIVE
	cell.rule(2)
	if cell.alive == DEAD {
		t.Errorf("Expected cell alive but it died.")
	}
}

func TestRule2AliveCellShouldBeStillAliveWhenThreeNeighboursHasAlive(t *testing.T) {
	var cell life
	cell.alive = LIVE
	cell.rule(3)
	if cell.alive == DEAD {
		t.Errorf("Expected cell alive but it died.")	
	}
}

func TestRule3AliveCellShouldBeDiedWhenMoreThanThreeNeighboursHasAlive(t *testing.T) {
	var cell life
	cell.alive = LIVE
	cell.rule(4)
	if cell.alive == LIVE {
		t.Errorf("Expected cell died but it alive.")	
	}
}

func TestRule4DeadCellShouldBeBornWhenThreeNeighboursHasAlive(t *testing.T) {
	var cell life
	cell.alive = DEAD
	cell.rule(3)
	if cell.alive == DEAD {
		t.Errorf("Expected cell alive but it died.")	
	}
}

func TestRule4DeadCellShouldBeStillDeadWhenTwoNeighboursHasAlive(t *testing.T) {
	var cell life
	cell.alive = DEAD
	cell.rule(2)
	if cell.alive == LIVE {
		t.Errorf("Expected cell died but it alive.")	
	}
}
