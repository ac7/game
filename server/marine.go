package server

func makeMarine() IUnit {
	return &unit{
		id:        newId(),
		name:      "Marine",
		health:    int64(70),
		maxHealth: int64(70),
		tags:      TAG_BIO | TAG_RANGED,
	}
}
