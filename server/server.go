package server

import (
	"math/rand"
)

const HANDSHAKE_CLIENT = "handshake"
const HANDSHAKE_SERVER = "handshake_part_two"

const (
	TAG_BIO = int64(1) << iota
	TAG_MECH
	TAG_FLYING
	TAG_CLOAKED
	TAG_SHIELDED
	TAG_RANGED
)

func newId() int64 {
	return rand.Int63()
}

type IUnit interface {
	Id() int64
	Name() string
	Health() int64
	MaxHealth() int64
	Tags() int64
	Position() (float64, float64)

	HasTags(tags ...int64) bool
	AddTags(tags ...int64)
	RemoveTags(tags ...int64)
	TakeDamage(amount int64) (alive bool)
	SetPosition(float64, float64)
	Distance(other IUnit) float64
}

type IWorld interface {
	AddUnit(unit IUnit)
	Unit(id int64) IUnit
	Units() []IUnit
	RemoveUnit(id int64) (success bool)
}
