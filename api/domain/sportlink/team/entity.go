package team

import "fmt"

type Category uint32

const (
	Unranked    Category = 0
	L1                   = 1
	L2                   = 2
	L3                   = 3
	L4                   = 4
	L5                   = 5
	L6                   = 6
	L7                   = 7
	MaxCategory          = L7
)

func GetCategory(value int) (Category, error) {
	if value < int(Unranked) || value > int(MaxCategory) {
		return 0, fmt.Errorf("invalid category value: %d", value)
	}
	return Category(value), nil
}

type Sport string

const (
	Paddle   Sport = "Paddle"
	Football Sport = "Football"
	Tennis   Sport = "Tennis"
)

type ID struct {
	Name  string
	Sport Sport
}
type Entity struct {
	ID       string
	Category Category
	Sport    Sport
}
