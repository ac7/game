package server

import (
	"testing"
)

func TestMarine(t *testing.T) {
	m := makeMarine()
	if !m.HasTags(TAG_BIO, TAG_RANGED) {
		t.Errorf("New Marine doesn't have BIO and RANGE tags.")
	}
	m.RemoveTags(TAG_RANGED)
	if m.HasTags(TAG_RANGED) {
		t.Errorf("After removing RANGE tag, HasTags(TAG_RANGED) still returend true.")
	}
	if m.HasTags(TAG_BIO, TAG_MECH) {
		t.Errorf("Marine is not supposed to have a TAG_MECH tag.")
	}

	initialHealth := m.Health()
	m.TakeDamage(3)
	if m.Health() != initialHealth-3 {
		t.Errorf("Unit did not decrease to %d health as expected (health was %d)",
			initialHealth-3, m.Health())
	}

	if _, ok := m.(*Marine); !ok {
		t.Errorf("Could not cast to Marine.")
	}
}
