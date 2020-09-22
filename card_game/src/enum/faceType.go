package enum

type FaceType int

const (
	Ace FaceType = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (d FaceType) String() string {
	return [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
		"Ten", "Jack", "Queen", "King"}[d-1]
}
