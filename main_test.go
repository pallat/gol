package main

import (
	"testing"
)

func TestRule1(t *testing.T) {
	var cell life
	cell.alive = true
	cell.rule1(1)
	if cell.alive == LIVE {
		t.Errorf("Expectec cell died but it alive %v", cell.alive)
	}
}