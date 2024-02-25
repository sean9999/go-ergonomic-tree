package ergotree

import (
	"errors"
	"slices"
)

type Node[K comparable] interface {
	Data() K
	Children() map[K]Node[K]
	Spawn(K) Node[K]
	RemoveChild(K)
	Parent() Node[K]
	Walk() [][]K
	Ancestry() []K
	IsTerminal() bool
	Set(K, Node[K]) error
	Get(K) (Node[K], bool)
}

// node implements Node
type node[K comparable] map[K]Node[K]

// Spawn creates and returns a child Node
func (t *node[K]) Spawn(key K) Node[K] {
	child := New(t)
	t.Set(key, child)
	return child
}

// IsTerminal returns true if the Node contains no child Nodes
func (t *node[K]) IsTerminal() bool {
	// if length of map is 1, that means it only contains a pointer to it's parent. Nothing else
	// length should never be zero, even for the root node, whose parent is nil
	lengthOfMap := len(*t)
	if lengthOfMap == 0 {
		panic("length of map should never be zero")
	}
	return len(*t) == 1
}

// // SetParent allows a Node to know who it's parent is
// func (t *node[K]) SetParent(parent Node[K]) {
// 	var zerok K
// 	(*t)[zerok] = parent
// }

// RemoveChild removes a child Node
func (t *node[K]) RemoveChild(key K) {
	delete(*t, key)
}

// basic equality check
func (t *node[K]) Equals(t2 Node[K]) bool {
	return t == t2.(*node[K])
}

// get key of self by querying parent
// zero value means root (no key)
func (t *node[K]) Data() K {
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
func (t *node[K]) Parent() Node[K] {
	var zerok K
	return (*t)[zerok]
}

// returns a slice of []K, indicating the path to the Node
func (t *node[K]) Ancestry() []K {
	var zerok K
	ancestry := []K{}
	ancestor := Node[K](t)
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
func (t *node[K]) Walk() [][]K {
	things := [][]K{}
	var acc func(Node[K])
	acc = func(tree Node[K]) {
		for _, subTree := range t.Children() {
			if subTree.IsTerminal() {
				things = append(things, subTree.Ancestry())
			} else {
				acc(subTree)
			}
		}
	}
	acc(t)
	return things
}

// Set sets a value on the map
func (t *node[K]) Set(key K, val Node[K]) error {
	var err error
	var zerok K
	if key == zerok {
		err = errors.New("cannot use the zero value as a key. That's reserved")
	} else {
		if !val.IsTerminal() {
			setParent[K](val.(*node[K]), t)
		}
		(*t)[key] = val
	}
	return err
}

// Get gets the value from the map
func (t *node[K]) Get(key K) (Node[K], bool) {
	val, ok := (*t)[key]
	return val, ok
}

// Children returns the map, omiting the special "parent" entry
func (t *node[K]) Children() map[K]Node[K] {
	var zerok K
	children := map[K]Node[K]{}
	for k, v := range *t {
		if k != zerok {
			children[k] = v
		}
	}
	return children
}

// New is a constructor
func New[K comparable](parent Node[K]) Node[K] {
	me := new(node[K])
	setParent[K](me, parent.(*node[K]))
	return me
}
