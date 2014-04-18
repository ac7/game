package server

import (
	"fmt"
	"strconv"
	"strings"
)

func Serialize(u IUnit) string {
	x, y := u.Position()
	return fmt.Sprintf(`%s;%d;%d;%d;%s;%s`, u.Name(), u.Health(),
		u.MaxHealth(), u.Tags(),
		strconv.FormatFloat(x, 'f', -1, 64),
		strconv.FormatFloat(y, 'f', -1, 64))
}

func Deserialize(s string) (IUnit, error) {
	u := new(Unit)
	var err error

	splits := strings.Split(s, ";")
	if len(splits) != 6 {
		return nil, fmt.Errorf("Splitting string yielded %d fragments (expected %d)",
			len(splits), 6)
	}

	u.name = splits[0]
	u.health, err = strconv.Atoi(splits[1])
	if err != nil {
		return nil, err
	}
	u.maxHealth, err = strconv.Atoi(splits[2])
	if err != nil {
		return nil, err
	}
	u.tags, err = strconv.Atoi(splits[3])
	if err != nil {
		return nil, err
	}
	u.x, err = strconv.ParseFloat(splits[4], 64)
	if err != nil {
		return nil, err
	}
	u.y, err = strconv.ParseFloat(splits[5], 64)
	if err != nil {
		return nil, err
	}
	return u, nil
}
