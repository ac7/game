package server

import (
	"math/rand"
)

const HANDSHAKE_CLIENT = "handshake"
const HANDSHAKE_SERVER = "handshake_part_two"

const (
	TAG_BIO = 1 << iota
	TAG_MECH
	TAG_FLYING
	TAG_CLOAKED
	TAG_SHIELDED
	TAG_RANGED
)

func newId() int {
	return rand.Int()
}

type IUnit interface {
	Id() int
	Name() string
	Health() int
	MaxHealth() int
	Tags() int
	Position() (float64, float64)

	HasTags(tags ...int) bool
	AddTags(tags ...int)
	RemoveTags(tags ...int)
	TakeDamage(amount int) (alive bool)
	SetPosition(float64, float64)
	Distance(other IUnit) float64
}

type IWorld interface {
	AddUnit(unit IUnit)
	Unit(id int) IUnit
	Units() []IUnit
	RemoveUnit(id int) (success bool)
}
