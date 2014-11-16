package main

const (
	LIVE = true
	DEAD = false
)

type life struct {
	alive bool
}

func (l *life)rule1(neighbours int) {
	l.alive = DEAD
}

func main() {
	
}

