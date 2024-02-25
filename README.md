# ergotree

ergotree is a Go package that implements a tree data-structure using a recursively defined map as it's backing store.

## Getting Started

```shell
go get github.com/sean9999/go-ergonomic-tree
```

## Example

```go

package main

import (
	"reflect"
	"testing"

	ergotree "github.com/sean9999/go-ergonomic-tree"
)

func TestErgonomicTree(t *testing.T) {

	//  create a tree
	life := ergotree.New[string](nil)

	//  create some children
	greatApes := life.Spawn("Primates").Spawn("Apes").Spawn("Great Apes")

	//  create some leaf nodes
	greatApes.Set("Western Gorilla")
	greatApes.Set("Eastern Gorilla")

	//  walk the whole tree, returning full paths to all leaf nodes
	got := life.Walk()

	//  check the data
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

```