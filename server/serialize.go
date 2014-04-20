package server

import (
	"log"

	"github.com/vmihailenco/msgpack"
)

func Serialize(u IUnit) string {
	x, y := u.Position()
	msg, err := msgpack.Marshal(map[string]interface{}{
		"id":    u.Id(),
		"name":  u.Name(),
		"hp":    u.Health(),
		"maxhp": u.MaxHealth(),
		"tags":  u.Tags(),
		"x":     x,
		"y":     y,
	})

	if err != nil {
		log.Fatal(err)
	}
	return string(msg)
}

// FIXME: this is dangerous, a bad packet would crash the server
// need an elegant way to check for presence in the map
func Deserialize(s string) (IUnit, error) {
	u := new(unit)

	var attrs map[string]interface{}
	err := msgpack.Unmarshal([]byte(s), &attrs)
	if err != nil {
		return nil, err
	}

	u.id = attrs["id"].(int64)
	u.name = attrs["name"].(string)
	u.health = attrs["hp"].(int64)
	u.maxHealth = attrs["maxhp"].(int64)
	u.tags = attrs["tags"].(int64)
	u.x = attrs["x"].(float64)
	u.y = attrs["y"].(float64)

	return u, nil
}
