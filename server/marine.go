package server

func makeMarine() IUnit {
	return &Unit{
		name:      "Marine",
		health:    70,
		maxHealth: 70,
		tags:      TAG_BIO | TAG_RANGED,
	}
}
