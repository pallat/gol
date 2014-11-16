package main

const (
	LIVE = true
	DEAD = false
)

type life struct {
	alive bool
}

func (l *life)rule(neighbours int) {
	if neighbours < 2 || neighbours > 3{
		l.alive = DEAD
	}else {
		l.alive = LIVE
	}
}

func main() {
	
}

