package server

import (
	"io"
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
	TAG_MELEE
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
	Serialize() string
	Deserialize(packet string) error

	HasTags(tags ...int64) bool
	AddTags(tags ...int64)
	RemoveTags(tags ...int64)
	TakeDamage(amount int64) (alive bool)
	SetPosition(float64, float64)
	Distance(other IUnit) float64
}

type IWorld interface {
	AddUnit(unit IUnit)
	RemoveUnit(id int64) (success bool)

	// Retrieve unit by id. Return nil for no such unit found.
	Unit(id int64) IUnit

	// Return a slice of all units in the world.
	Units() []IUnit
}

type IServer interface {
	// The location parameter should look like "localhost:1030"
	Listen(location string) error

	// HandleConn() is called automatically when a client connects
	// to the server, but it can be called manually for testing.
	HandleConn(conn io.ReadWriteCloser, id int) error
}
