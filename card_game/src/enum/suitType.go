package enum

type SuitType int

const (
	Club SuitType = iota + 1
	Diamond
	Spade
	Heart
)

func (d SuitType) String() string {
	return [...]string{"Club", "Diamond", "Spade", "Heart"}[d-1]
}
