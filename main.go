package main

import (
	// "fmt"
)

const (
	LIVE = true
	DEAD = false
)

type life struct {
	alive bool
}

func (l *life)rule(neighbours int) {
	if l.alive == DEAD {
		if neighbours == 3 {
			l.alive = LIVE
		}
	}else {
		if neighbours == 2 || neighbours == 3{
			l.alive = LIVE
		}else {
			l.alive = DEAD
		}
	}
}

type table [][]bool

func (t *table) born(x, y int) {
	[][]bool(*t)[x][y] = true
}

func(t *table) around(x, y int) int {
	xmin := 0
	ymin := 0

	count :=0
	if x > 0 {xmin = x-1}
	if y > 0 {ymin = y-1}

	for xx:=xmin;xx<=x+1;xx++ {
		for yy:=ymin;yy<=y+1;yy++ {
			if 	[][]bool(*t)[xx][yy] != true {
				count +=1
			}

		}
	}

	return count
}

func spaces(height int, weight int) table {
	var t table
	for i:=0;i<height;i++ {
		t = append(t,make([]bool,weight))
	}

	return t
}

func main() {
	
}

