package server

type Marine struct {
	Unit
}

func makeMarine() IUnit {
	return &Marine{
		Unit{
			name:      "Marine",
			health:    70,
			maxHealth: 70,
			tags:      TAG_BIO | TAG_RANGED,
		},
	}
}

func (m *Marine) Rank() string {
	return "Private"
}
