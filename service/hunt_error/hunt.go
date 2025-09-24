package hunt_error

type Hunt struct {
	Cat Cat
}

func NewHunt(cat Cat) Hunt {
	return Hunt{cat}
}
