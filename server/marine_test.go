package server

import (
	"math"
	"testing"
)

func TestMarine(t *testing.T) {
	m := makeMarine()

	if x, y := m.Position(); x != 0.0 || y != 0.0 {
		t.Errorf("Marine initial position is not 0, 0")
	}

	if !m.HasTags(TAG_BIO, TAG_RANGED) {
		t.Errorf("New Marine doesn't have BIO and RANGE tags.")
	}

	if m.RemoveTags(TAG_RANGED); m.HasTags(TAG_RANGED) {
		t.Errorf("After removing RANGE tag, HasTags(TAG_RANGED) still returend true.")
	}
	if m.HasTags(TAG_BIO, TAG_MECH) {
		t.Errorf("Marine is not supposed to have a TAG_MECH tag.")
	}

	initialHealth := m.Health()
	if m.TakeDamage(3); m.Health() != initialHealth-3 {
		t.Errorf("Unit did not decrease to %d health as expected (health was %d)",
			initialHealth-3, m.Health())
	}

	if _, ok := m.(*Marine); !ok {
		t.Errorf("Could not cast to *Marine")
	}
}

func TestDistance(t *testing.T) {
	m1 := makeMarine()
	m2 := makeMarine()

	m1.SetPosition(0.0, 0.0)
	m2.SetPosition(0.0, 5.0)
	if m1.Distance(m2) != 5.0 || m2.Distance(m1) != 5.0 {
		t.Errorf("Unit.Distance() returned incorrect value %f (expected %f)",
			m1.Distance(m2), 5.0)
	}

	m1.SetPosition(2.5, 2.0)
	m2.SetPosition(4.0, 6.0)
	if expected := math.Hypot(1.5, 4.0); m1.Distance(m2) != expected {
		t.Errorf("Unit.Distance() returned incorrect value %f (expected %f)",
			m1.Distance(m2), expected)
	}
}
