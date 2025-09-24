package hunt_error

import (
	"errors"
	"time"
)

type Cat struct {
	Name string
}

func NewCat(name string) (Cat, error) {
	if time.Now().Unix()%2 == 0 {
		return Cat{}, errors.New("cat dead")
	}
	return Cat{Name: name}, nil
}
