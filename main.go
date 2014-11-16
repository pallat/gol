package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	LIVE = true
	DEAD = false
)

type life struct {
	alive bool
}

func (l *life)rule(neighbours int) {
// 	LIVE 1 DEAD
// >	LIVE 2 LIVE
// >	LIVE 3 LIVE
// 	LIVE 4 DEAD
// 	LIVE 5 DEAD
// 	LIVE 6 DEAD
// 	LIVE 7 DEAD
// 	LIVE 8 DEAD
// 	DEAD 1 DEAD
// 	DEAD 2 DEAD
// >	DEAD 3 LIVE
// 	DEAD 4 DEAD
// 	DEAD 5 DEAD
// 	DEAD 6 DEAD
// 	DEAD 7 DEAD
// 	DEAD 8 DEAD

	if neighbours == 3 {
		l.alive = LIVE
	}else if neighbours == 2 && l.alive == LIVE{
		l.alive = LIVE
	}else {
		l.alive = DEAD
	}

}

type table [][]bool

func (t *table) born(x, y int) {
	[][]bool(*t)[y][x] = true
}

func(t *table) around(x, y int) int {
	xmin := 0
	ymin := 0

	count :=0
	if x > 0 {xmin = x-1}
	if y > 0 {ymin = y-1}

	for xx:=xmin;xx<=x+1;xx++ {
		for yy:=ymin;yy<=y+1;yy++ {
			if 	[][]bool(*t)[yy][xx] != true {
				count +=1
			}

		}
	}

	return count
}

func(t *table) neighbours(x, y int) int {
	xmin := 0
	ymin := 0
	xmax := x+1
	ymax := y+1


	count :=0
	if x > 0 {xmin = x-1}
	if y > 0 {ymin = y-1}

	tab := [][]bool(*t)
	if xmax >= len(tab[0]) {
		xmax = len(tab[0])-1
	}
	if ymax >= len(tab) {
		ymax = len(tab)-1
	}

	for yy:=ymin; yy<=ymax; yy++ {
		for xx:=xmin; xx<=xmax; xx++ {
			if 	[][]bool(*t)[yy][xx] == true {
				// fmt.Printf("x %v, y %v, value %v\n",xx,yy,[][]bool(*t)[yy][xx])
				count +=1
			}

		}
	}

	if count > 0 {
		count -= 1
	}
	return count
}

func (t *table) turn() {
	tab := [][]bool(*t)
	newTable := spaces(len(tab),len(tab[0]))
	for y := range tab {
		for x := range tab[y] {
			neighbours := t.neighbours(x,y)

			cell := life{
				alive: tab[y][x],
			}
			cell.rule(neighbours)
			[][]bool(newTable)[y][x] = cell.alive
		}
	}

	*t = newTable
}

func spaces(weight, height int) table {
	var t table
	for i:=0;i<height;i++ {
		t = append(t,make([]bool,weight))
	}

	return t
}

func main() {
	begin := 40
	spaces := spaces(begin,begin)
	rand.Seed(time.Now().UnixNano())
	total := rand.Intn(begin * 10)

	for i:=0; i<total; i++ {
		spaces.born(rand.Intn(begin),rand.Intn(begin))		
	}

	tab := [][]bool(spaces)

	for {

    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()

		for y := range tab {
			for x := range tab[y] {
				if [][]bool(spaces)[y][x] {
					fmt.Printf("O")
				}else {
					fmt.Printf("_")
				}
			}
			fmt.Printf("\n")
		}

		spaces.turn()

		time.Sleep(1000 * time.Millisecond)
	}
}

