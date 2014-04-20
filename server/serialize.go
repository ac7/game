package server

import (
	"log"

	"github.com/vmihailenco/msgpack"
)

func Serialize(u IUnit) string {
	x, y := u.Position()
	msg, err := msgpack.Marshal(map[string]interface{}{
		"name":  u.Name(),
		"hp":    int64(u.Health()),
		"maxhp": int64(u.MaxHealth()),
		"tags":  int64(u.Tags()),
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
	u := new(Unit)

	var attrs map[string]interface{}
	err := msgpack.Unmarshal([]byte(s), &attrs)
	if err != nil {
		return nil, err
	}

	u.name = attrs["name"].(string)
	u.health = int(attrs["hp"].(int64))
	u.maxHealth = int(attrs["maxhp"].(int64))
	u.tags = int(attrs["tags"].(int64))
	u.x = attrs["x"].(float64)
	u.y = attrs["y"].(float64)

	return u, nil
}
