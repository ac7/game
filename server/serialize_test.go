package server

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	cases := []struct {
		unit IUnit
		str  string
	}{
		{
			&Unit{"Marine", 70, 70, 0, 32.0, 18.0},
			"Marine;70;70;0;32;18",
		},
		{
			&Unit{
				"random string fragment", 18, 23,
				TAG_BIO & TAG_MECH, -3.4, -111.00,
			},
			fmt.Sprintf("random string fragment;18;23;%d;-3.4;-111",
				TAG_BIO&TAG_MECH,
			),
		},
	}
	for _, c := range cases {
		deserializedUnit, err := Deserialize(c.str)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(deserializedUnit.(*Unit), c.unit) {
			t.Errorf("Didn't deserialize %s correctly.", c.str)
		}

		packet := Serialize(c.unit)
		if packet != c.str {
			t.Errorf("Serializing the unit yielded '%s' (expected '%s')",
				packet, c.str)
		}
	}
}
