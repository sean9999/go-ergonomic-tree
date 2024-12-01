package ergotree_test

import (
	"slices"
	"testing"

	tree "github.com/sean9999/go-ergonomic-tree"
)

func TestNew(t *testing.T) {

	//	life is the root node of the tree of life
	life := tree.New[string](nil)

	//	Great Apes descend from Apes, which descend from Primates
	greatApes := life.Spawn("Primates").Spawn("Apes").Spawn("Great Apes")

	//	Eastern and Western Gorilla are each distinct species which descend from Great Apes
	greatApes.Spawn("Eastern Gorilla")
	greatApes.Spawn("Western Gorilla")

	//	walk the tree of life (depth first)
	got := life.Walk()

	want := [][]string{
		{"Primates", "Apes", "Great Apes", "Eastern Gorilla"},
		{"Primates", "Apes", "Great Apes", "Western Gorilla"},
	}

	//	we cannot be guratantted in which order we get trees of the branch
	//	since our backing data structure is a map
	if slices.Compare(want[0], got[0]) != 0 && slices.Compare(want[1], got[0]) != 0 {
		t.Errorf("got %v but wanted %v", got, want)
	}
	if slices.Compare(want[0], got[1]) != 0 && slices.Compare(want[1], got[1]) != 0 {
		t.Errorf("got %v but wanted %v", got, want)
	}

}
