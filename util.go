package ergotree

func setParent[K comparable](child *node[K], parent *node[K]) {
	var zerok K
	(*child)[zerok] = parent
}
