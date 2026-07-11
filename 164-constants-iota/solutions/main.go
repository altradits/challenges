package piscine

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func TurnRight(d Direction) Direction {
	return (d + 1) % 4
}

func TurnLeft(d Direction) Direction {
	return (d + 3) % 4
}

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type Permission uint

const (
	Read    Permission = 1 << iota
	Write
	Execute
)

func HasPermission(p, flag Permission) bool {
	return p&flag != 0
}
