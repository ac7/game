package server

const (
	TAG_BIO = 1 << iota
	TAG_MECH
	TAG_FLYING
	TAG_CLOAKED
	TAG_SHIELDED
	TAG_RANGED
)

type IUnit interface {
	Name() string
	Health() int
	MaxHealth() int
	Tags() int

	HasTags(tags ...int) bool
	AddTags(tags ...int)
	RemoveTags(tags ...int)
	TakeDamage(amount int) (alive bool)
	SetPosition(float64, float64)
	Position() (float64, float64)
	Distance(other IUnit) float64
}

func isBio(unit IUnit) bool {
	return (unit.Tags()&TAG_BIO > 0)
}
