package server

import (
	"math"
)

type unit struct {
	id        int
	name      string
	health    int
	maxHealth int
	tags      int

	x float64
	y float64
}

// getters
func (u *unit) Id() int        { return u.id }
func (u *unit) Name() string   { return u.name }
func (u *unit) Health() int    { return u.health }
func (u *unit) MaxHealth() int { return u.maxHealth }
func (u *unit) Tags() int      { return u.tags }

// functionality
func (u *unit) AddTags(tags ...int) {
	for _, tag := range tags {
		u.tags &= tag
	}
}

func (u *unit) RemoveTags(tags ...int) {
	for _, tag := range tags {
		u.tags ^= tag
	}
}

func (u *unit) HasTags(tags ...int) bool {
	for _, tag := range tags {
		if u.tags&tag == 0 {
			return false
		}
	}
	return true
}

func (u *unit) SetPosition(x, y float64) {
	u.x = x
	u.y = y
}

func (u *unit) Position() (float64, float64) {
	return u.x, u.y
}

func (u *unit) Distance(other IUnit) float64 {
	otherX, otherY := other.Position()
	return math.Hypot(u.x-otherX, u.y-otherY)
}

func (u *unit) TakeDamage(amount int) (alive bool) {
	u.health -= amount
	alive = u.health > 0
	return
}
