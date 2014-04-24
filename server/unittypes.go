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

func makeMedic() IUnit {
	return &unit{
		id:        newId(),
		name:      "Medic",
		health:    int64(50),
		maxHealth: int64(50),
		tags:      TAG_BIO,
	}
}
