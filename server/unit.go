package server

import (
	"math"
)

type unit struct {
	id        int64
	name      string
	health    int64
	maxHealth int64
	tags      int64

	x float64
	y float64
}

// getters
func (u *unit) Id() int64        { return u.id }
func (u *unit) Name() string     { return u.name }
func (u *unit) Health() int64    { return u.health }
func (u *unit) MaxHealth() int64 { return u.maxHealth }
func (u *unit) Tags() int64      { return u.tags }

// functionality
func (u *unit) AddTags(tags ...int64) {
	for _, tag := range tags {
		u.tags &= tag
	}
}

func (u *unit) RemoveTags(tags ...int64) {
	for _, tag := range tags {
		u.tags ^= tag
	}
}

func (u *unit) HasTags(tags ...int64) bool {
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

func (u *unit) TakeDamage(amount int64) (alive bool) {
	u.health -= amount
	alive = u.health > 0
	return
}
