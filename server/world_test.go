package server

import (
	"testing"
)

func TestWorld(t *testing.T) {
	w := newWorld().(*world)
	m := makeMarine()

	w.AddUnit(makeMarine())
	w.AddUnit(makeMarine())
	w.AddUnit(m)
	w.AddUnit(makeMarine())

	if m != w.Unit(m.Id()) {
		t.Errorf("Could not retrieve unit by id.")
	}
	if !w.RemoveUnit(m.Id()) {
		t.Errorf("Could not remove unit by id.")
	}
	if w.Unit(m.Id()) != nil {
		t.Errorf("Could fetch unit after removing it (shouldn't be possible).")
	}
	if len(w.Units()) != 3 {
		t.Errorf("After removing unit, count was %d (expected %d)",
			len(w.Units()), 3)
	}
	for i := 0; i < 3; i++ {
		if !w.RemoveUnit(w.Units()[0].Id()) {
			t.Errorf("Couldn't remove unit from list.")
		}
	}

	if w.RemoveUnit(0xdeadbeef) {
		t.Errorf("RemoveUnit() worked for invalid unit.")
	}
}
