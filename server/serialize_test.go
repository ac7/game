package server

import (
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	cases := []IUnit{
		&Unit{newId(), "Marine", 70, 70, 0, 32.0, 18.0},
		&Unit{
			newId(),
			"random string fragment", 18, 23,
			TAG_BIO & TAG_MECH, -3.4, -111.00,
		},
	}

	for _, unit := range cases {
		t.Logf("Testing serialization for unit %+v", unit)
		str := Serialize(unit)
		deserializedUnit, err := Deserialize(str)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(deserializedUnit.(*Unit), unit) {
			t.Errorf("Didn't deserialize %s correctly.", str)
		}
	}
}
