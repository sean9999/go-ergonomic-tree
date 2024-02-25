package ergotree_test

import (
	"reflect"
	"testing"

	ergotree "github.com/sean9999/go-ergonomic-tree"
)

func TestNew(t *testing.T) {

	life := ergotree.New[string](nil)

	primates := life.Spawn("Primates")
	apes := primates.Spawn("Apes")
	greatApes := apes.Spawn("Great Apes")
	greatApes.Set("Eastern Gorilla")
	//greatApes.Set("Western Gorilla")

	got := life.Walk()

	want := [][]string{
		{"Primates", "Apes", "Great Apes", "Eastern Gorilla"},
		//{"Primates", "Apes", "Great Apes", "Western Gorilla"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}

}
