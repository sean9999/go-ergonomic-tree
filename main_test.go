package ergotree_test

import (
	"reflect"
	"testing"

	ergotree "github.com/sean9999/go-ergonomic-tree"
)

func TestNew(t *testing.T) {

	life := ergotree.New[string](nil)

	greatApes := life.Spawn("Primates").Spawn("Apes").Spawn("Great Apes")

	greatApes.Set("Eastern Gorilla")
	greatApes.Set("Western Gorilla")

	got := life.Walk()

	want := [][]string{
		{"Primates", "Apes", "Great Apes", "Eastern Gorilla"},
		{"Primates", "Apes", "Great Apes", "Western Gorilla"},
	}

	if !(reflect.DeepEqual(got[0], want[0]) || reflect.DeepEqual(got[0], want[1])) {
		t.Errorf("got %v but wanted %v", got, want)
	}

	if !(reflect.DeepEqual(got[1], want[0]) || reflect.DeepEqual(got[1], want[1])) {
		t.Errorf("got %v but wanted %v", got, want)
	}

}
