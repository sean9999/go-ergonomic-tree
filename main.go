package ergotree

import (
	"encoding/json"
	"slices"
)

type Node[K comparable] map[K]*Node[K]

//	a Node is a Tree is a Node
// GOEXPERIMENT=aliastypeparams
//type Tree[K comparable] = Node[K]

func (t *Node[K]) SetParent(p *Node[K]) {
	var zerok K
	(*t)[zerok] = p
}

// IsTerminal returns true if the Node contains no child Nodes
func (t *Node[K]) IsTerminal() bool {

	// if length of map is 1, that means it only contains a pointer to it's parent. Nothing else
	// length should never be zero, even for the root node, whose parent is nil
	lengthOfMap := len(*t)
	if lengthOfMap == 0 {
		panic("length of map should never be zero")
	}
	return lengthOfMap == 1
}

// RemoveChild removes a child Node
func (t *Node[K]) RemoveChild(key K) {
	delete(*t, key)
}

// basic equality check
func (t *Node[K]) Equals(t2 *Node[K]) bool {
	return t == t2
}

// get key of self by querying parent
// zero value means root (no key)
func (t *Node[K]) Data() K {
	var magicKey K
	parent := t.Parent()
	if parent != nil {
		for thisKey, sibling := range parent.Children() {
			if sibling == t {
				magicKey = thisKey
			}
		}
	}
	return magicKey
}

// Parent returns a Node's parent
func (t *Node[K]) Parent() *Node[K] {
	var zerok K
	return (*t)[zerok]
}

// returns a slice of []K, indicating the path to the Node
func (t *Node[K]) Ancestry() []K {
	var zerok K
	ancestry := []K{}
	ancestor := t
	for {
		if ancestor == nil {
			break
		}
		data := ancestor.Data()
		if data == zerok {
			break
		}
		ancestry = append(ancestry, data)
		ancestor = ancestor.Parent()
	}
	slices.Reverse(ancestry)
	return ancestry
}

// Walk returns all leaf nodes as ancestries
func (t *Node[K]) Walk() [][]K {
	things := [][]K{}
	var acc func(*Node[K])

	acc = func(tree *Node[K]) {
		for _, subTree := range tree.Children() {

			switch {
			case subTree == nil:
				// do nothing
			case subTree.IsTerminal():
				things = append(things, subTree.Ancestry())
			default:
				acc(subTree)
			}

		}
	}

	acc(t)
	return things
}

func (t *Node[K]) String() string {
	j, err := json.Marshal(t.Walk())
	if err != nil {
		panic(err)
	}
	return string(j)
}

// Spawn creates and returns a child Node
func (t *Node[K]) Spawn(key K) *Node[K] {
	child := New(t)
	(*t)[key] = child
	return child
}

// Set simply calls Spawn(), but discards return value
func (t *Node[K]) Set(key K) {
	t.Spawn(key)
}

// Get gets the value from the map
func (t *Node[K]) Get(key K) (*Node[K], bool) {
	val, ok := (*t)[key]
	return val, ok
}

// Children returns the map, omiting the special "parent" entry
func (t *Node[K]) Children() map[K]*Node[K] {
	var zerok K
	children := map[K]*Node[K]{}
	for k, v := range *t {
		if k != zerok {
			children[k] = v
		}
	}
	return children
}

// New is a constructor
func New[K comparable](parent *Node[K]) *Node[K] {
	me := &Node[K]{}
	me.SetParent(parent)
	return me
}
