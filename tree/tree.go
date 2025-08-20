package tree

import "Goinator/data"

type Node struct {
	Question string
	Yes      *Node
	No       *Node
	Entity   *data.Entity // only filled if it’s a leaf
}

// Hard-coded tree for now
var Root = &Node{
	Question: "Is it a real person?",
	Yes: &Node{
		Question: "Is it male?",
		Yes: &Node{
			Question: "Is he famous for science?",
			Yes:      &Node{Entity: &data.Entity{Name: "Albert Einstein"}},
			No:       &Node{Entity: &data.Entity{Name: "Tom Cruise"}},
		},
		No: &Node{
			Question: "Is she an actress?",
			Yes:      &Node{Entity: &data.Entity{Name: "Scarlett Johansson"}},
			No:       &Node{Entity: &data.Entity{Name: "Marie Curie"}},
		},
	},
	No: &Node{
		Question: "Is it a cartoon?",
		Yes:      &Node{Entity: &data.Entity{Name: "Mickey Mouse"}},
		No:       &Node{Entity: &data.Entity{Name: "Rock"}},
	},
}
