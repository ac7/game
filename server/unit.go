package server

import (
	"math"
)

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

type Unit struct {
	name      string
	health    int
	maxHealth int
	tags      int

	x float64
	y float64
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

func (u *Unit) SetPosition(x, y float64) {
	u.x = x
	u.y = y
}

func (u *Unit) Position() (float64, float64) {
	return u.x, u.y
}

func (u *Unit) Distance(other IUnit) float64 {
	otherX, otherY := other.Position()
	return math.Hypot(u.x-otherX, u.y-otherY)
}

func (u *Unit) TakeDamage(amount int) (alive bool) {
	u.health -= amount
	alive = u.health > 0
	return
}
