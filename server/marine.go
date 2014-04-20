package server

func makeMarine() IUnit {
	return &unit{
		name:      "Marine",
		health:    70,
		maxHealth: 70,
		tags:      TAG_BIO | TAG_RANGED,
	}
}
