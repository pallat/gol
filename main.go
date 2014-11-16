package main

const (
	LIVE = true
	DEAD = false
)

type life struct {
	alive bool
}

func (l *life)rule1(neighbours int) {
	if neighbours < 2 {
		l.alive = DEAD
	}else {
		l.alive = LIVE
	}
}

func main() {
	
}

