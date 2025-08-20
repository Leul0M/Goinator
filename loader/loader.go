package loader

import (
	_ "embed"
	"encoding/json"
	"os"

	"Goinator/data"
	"Goinator/tree"
)

// Record matches your renamed JSON.
type Record struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	Traits map[string]interface{} `json:"traits"` // bool | null
}

// LoadEntities reads data/entities.json into a slice of Record.
func LoadEntities(path string) ([]Record, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var recs []Record
	return recs, json.Unmarshal(b, &recs)
}

// BuildTree turns the flat list into a *tree.Node.
func BuildTree(recs []Record) *tree.Node {
	// stable order – adjust as you like
	keys := []string{
    "can_fly",
    "has_fur_feathers",
    "has_screen",
    "has_superpowers",
    "is_actor",
    "is_electronic",
    "is_famous",
    "is_footwear",
    "is_human",
    "is_large_country",
    "is_living",
    "is_male",
    "is_multiplayer",
    "is_musician",
    "is_personal",
    "is_photo_sharing",
    "is_politician",
    "is_portable",
    "is_real",
    "is_superhero",
    "is_sweet",
    "is_youtuber",
    "lives_in_water",
}
	return build(keys, recs)
}

// ---------- internal helpers ----------

// build recurses down the question list.
func build(keys []string, recs []Record) *tree.Node {
	if len(keys) == 0 || len(recs) == 1 {
		return &tree.Node{Entity: &data.Entity{Name: recs[0].Name}}
	}

	q := keys[0]
	yes, no, other := split(recs, q)

	n := &tree.Node{Question: q}
	if len(yes) > 0 {
		n.Yes = build(keys[1:], yes)
	}
	if len(no) > 0 {
		n.No = build(keys[1:], no)
	}
	if len(other) > 0 { // treat unknown/null as "no"
		n.No = build(keys[1:], append(no, other...))
	}
	return n
}

// split groups records by the value of a single trait.
func split(recs []Record, key string) (yes, no, other []Record) {
	for _, r := range recs {
		switch truth(r.Traits[key]) {
		case 1:
			yes = append(yes, r)
		case -1:
			no = append(no, r)
		default:
			other = append(other, r)
		}
	}
	return
}

// truth converts bool|null into our three-way int.
func truth(v interface{}) int8 {
	if v == nil {
		return 0
	}
	if b, ok := v.(bool); ok && b {
		return 1
	}
	return -1
}