package loader

import (
	"encoding/json"
	"math"
	"os"
	"sort"

	"Goinator/data"
	"Goinator/tree"
)

type Record struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	Traits map[string]interface{} `json:"traits"`
}

func LoadEntities(path string) ([]Record, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var recs []Record
	return recs, json.Unmarshal(b, &recs)
}

func BuildTree(recs []Record) *tree.Node {
	allKeys := keysFromRecords(recs)
	return build(allKeys, recs)
}

/* ---------- tree construction ---------- */

// build : recursive ID3 + Bayesian tie-breaker
func build(keys []string, recs []Record) *tree.Node {
	// 1) single entity → leaf
	if len(recs) == 1 {
		return &tree.Node{Entity: &data.Entity{Name: recs[0].Name}}
	}
	// 2) no more keys or no useful split → pick the highest-probability entity
	if len(keys) == 0 {
		return pickBestEntity(recs)
	}

	q := bestQuestion(recs, keys)
	if q == "" {
		return pickBestEntity(recs)
	}

	yes, no, _ := split(recs, q)
	nextKeys := remove(keys, q)

	node := &tree.Node{Question: q}
	if len(yes) > 0 {
		node.Yes = build(nextKeys, yes)
	}
	if len(no) > 0 {
		node.No = build(nextKeys, no)
	}
	return node
}

/* ---------- Bayesian tie-breaker ---------- */

// pickBestEntity returns a leaf with the single highest-probability entity.
func pickBestEntity(recs []Record) *tree.Node {
	// uniform prior
	prior := 1.0 / float64(len(recs))
	bestName := recs[0].Name
	bestProb := 0.0

	for _, r := range recs {
		ll := math.Log(prior)
		for _, v := range r.Traits {
			switch truth(v) {
			case 1:
				ll += math.Log(0.9) // P(true | feature)
			case -1:
				ll += math.Log(0.1) // P(false | feature)
			}
		}
		prob := math.Exp(ll)
		if prob > bestProb {
			bestProb, bestName = prob, r.Name
		}
	}
	return &tree.Node{Entity: &data.Entity{Name: bestName}}
}

/* ---------- ID3 helpers ---------- */

func bestQuestion(recs []Record, keys []string) string {
	bestGain, bestKey := -1.0, ""
	total := float64(len(recs))
_ = total // silence unused (if the compiler complains)
	for _, k := range keys {
		yes, no, _ := split(recs, k)
		if len(yes) == 0 || len(no) == 0 {
			continue
		}
		gain := entropyGain(recs, yes, no)
		if gain > bestGain {
			bestGain, bestKey = gain, k
		}
	}
	return bestKey
}

func entropyGain(parent, yes, no []Record) float64 {
	total := float64(len(parent))
	if total == 0 {
		return 0
	}
	return entropy(parent) -
		(float64(len(yes))/total*entropy(yes) + float64(len(no))/total*entropy(no))
}

func entropy(recs []Record) float64 {
	m := make(map[string]int)
	for _, r := range recs {
		m[r.Name]++
	}
	total := float64(len(recs))
	h := 0.0
	for _, cnt := range m {
		p := float64(cnt) / total
		h -= p * math.Log2(p)
	}
	return h
}

/* ---------- utility ---------- */

func keysFromRecords(recs []Record) []string {
	m := make(map[string]struct{})
	for _, r := range recs {
		for k := range r.Traits {
			m[k] = struct{}{}
		}
	}
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func remove(s []string, x string) []string {
	r := make([]string, 0, len(s)-1)
	for _, v := range s {
		if v != x {
			r = append(r, v)
		}
	}
	return r
}

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

func truth(v interface{}) int8 {
	if v == nil {
		return 0
	}
	if b, ok := v.(bool); ok && b {
		return 1
	}
	return -1
}
