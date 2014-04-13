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
	AddTags(tags ...int)
	RemoveTags(tags ...int)
	HasTags(tags ...int) bool

	TakeDamage(amount int) (alive bool)
}

func isBio(unit IUnit) bool {
	return (unit.Tags()&TAG_BIO > 0)
}

type Unit struct {
	name      string
	health    int
	maxHealth int
	tags      int
}

// getters
func (u *Unit) Name() string   { return u.name }
func (u *Unit) Health() int    { return u.health }
func (u *Unit) MaxHealth() int { return u.maxHealth }
func (u *Unit) Tags() int      { return u.tags }

// functionality
func (u *Unit) AddTags(tags ...int) {
	for _, tag := range tags {
		u.tags &= tag
	}
}

func (u *Unit) RemoveTags(tags ...int) {
	for _, tag := range tags {
		u.tags ^= tag
	}
}

func (u *Unit) HasTags(tags ...int) bool {
	for _, tag := range tags {
		if u.tags&tag == 0 {
			return false
		}
	}
	return true
}

func (u *Unit) TakeDamage(amount int) (alive bool) {
	u.health -= amount
	alive = u.health > 0
	return
}
