package loader

import (
	"encoding/json"
	"math"
	"os"

	"Goinator/data"
	"Goinator/tree"
)

// Record matches the JSON file.
type Record struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	Traits map[string]interface{} `json:"traits"`
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

// BuildTree builds an ID3-style decision tree from the records.
func BuildTree(recs []Record) *tree.Node {
	// list of all possible trait keys
	allKeys := []string{
		"can_fly", "has_fur_feathers", "has_screen", "has_superpowers",
		"is_actor", "is_electronic", "is_famous", "is_footwear",
		"is_human", "is_large_country", "is_living", "is_male",
		"is_multiplayer", "is_musician", "is_personal", "is_photo_sharing",
		"is_politician", "is_portable", "is_real", "is_superhero",
		"is_sweet", "is_youtuber", "lives_in_water",
	}
	return build(allKeys, recs)
}

/* ---------- internal helpers ---------- */

// build constructs the tree recursively.
func build(keys []string, recs []Record) *tree.Node {
	// leaf if no more questions or only one entity left
	if len(keys) == 0 || len(recs) == 1 {
		return &tree.Node{Entity: &data.Entity{Name: recs[0].Name}}
	}

	// choose the best remaining question
	q := bestQuestion(recs, keys)
	if q == "" {
		// no useful split → leaf with the first entity
		return &tree.Node{Entity: &data.Entity{Name: recs[0].Name}}
	}

	// split records
	yes, no, _ := split(recs, q)

	// build nextKeys (q removed)
	nextKeys := make([]string, 0, len(keys)-1)
	for _, k := range keys {
		if k != q {
			nextKeys = append(nextKeys, k)
		}
	}

	node := &tree.Node{Question: q}
	if len(yes) > 0 {
		node.Yes = build(nextKeys, yes)
	}
	if len(no) > 0 {
		node.No = build(nextKeys, no)
	}
	return node
}

// bestQuestion returns the key with the highest information gain.
func bestQuestion(recs []Record, keys []string) string {
	bestGain, bestKey := -1.0, ""
	total := float64(len(recs))
	for _, k := range keys {
		yes, no, _ := split(recs, k)
		if len(yes) == 0 || len(no) == 0 {
			continue // no information
		}
		gain := entropy(recs) -
			(float64(len(yes))/total*entropy(yes) + float64(len(no))/total*entropy(no))
		if gain > bestGain {
			bestGain, bestKey = gain, k
		}
	}
	return bestKey
}

// split groups records by trait value.
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

// truth converts interface{} to ‑1,0,1.
func truth(v interface{}) int8 {
	if v == nil {
		return 0
	}
	if b, ok := v.(bool); ok && b {
		return 1
	}
	return -1
}

// entropy computes Shannon entropy of a set of records.
func entropy(recs []Record) float64 {
	n := float64(len(recs))
	if n == 0 {
		return 0
	}
	p := 1.0 // only one class per leaf in our current setup
	return -p * math.Log2(p)
}
