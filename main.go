package main

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

func main() {
	
}

