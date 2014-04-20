package server

type world struct {
	units []IUnit
}

// Returns nil if unit doesn't exist
func (w *world) Unit(id int) IUnit {
	index := w.unitIndex(id)
	if index < 0 {
		return nil
	}
	return w.units[index]
}

func (w *world) Units() []IUnit {
	return w.units
}

func (w *world) unitIndex(id int) int {
	for index, unit := range w.units {
		if unit.Id() == id {
			return index
		}
	}
	return -1
}

func (w *world) AddUnit(unit IUnit) {
	w.units = append(w.units, unit)
}

func (w *world) RemoveUnit(id int) (success bool) {
	index := w.unitIndex(id)
	if index < 0 {
		return false
	}

	w.units[index] = w.units[len(w.units)-1]
	w.units = w.units[0 : len(w.units)-1]
	return true
}

func newWorld() IWorld {
	return &world{
		units: []IUnit{},
	}
}
