package theme

import "sort"

const MainMenu = "Main"

var menuRegistry = make([]map[string]DataMenuItem, 0)

func RegisterMenu(m map[string]DataMenuItem) {
	menuRegistry = append(menuRegistry, m)
}

// sortMenu recursively sorts the menu tree in-place by the Sort field.
// For each DataMenuItem, its Children slice is sorted ascending by Sort,
// then the same is applied to each child.
func sortMenu(items []DataMenuItem) {
	// First, sort this level
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Sort < items[j].Sort
	})

	// Then recurse into children
	for idx := range items {
		if len(items[idx].Children) > 0 {
			sortMenu(items[idx].Children)
		}
	}
}

// mergeMenu combines multiple menu maps into one
//
// When top‑level keys coincide, their DataMenuItems are merged as follows:
// 1) Name, Href, and Attr are taken from the first encountered item, if they’re non‑empty.
// 2) Children are merged by Name: the order is preserved—first all unique items from the first slice, then from the second;
// if a Name appears in both, those entries are merged recursively
func mergeMenu(maps ...map[string]DataMenuItem) map[string]DataMenuItem {
	result := make(map[string]DataMenuItem)
	for _, m := range maps {
		for key, item := range m {
			if existing, ok := result[key]; ok {
				result[key] = mergeItem(existing, item)
			} else {
				result[key] = copyItem(item)
			}
		}
	}
	return result
}

func mergeItem(a, b DataMenuItem) DataMenuItem {
	out := copyItem(a)

	if out.Name == "" {
		out.Name = b.Name
	}
	if out.Href == "" {
		out.Href = b.Href
	}
	if out.Attr == "" {
		out.Attr = b.Attr
	}

	out.Children = mergeChildren(out.Children, b.Children)
	return out
}

func mergeChildren(a, b []DataMenuItem) []DataMenuItem {
	index := make(map[string]int, len(a))
	for i, it := range a {
		index[it.Name] = i
	}

	out := make([]DataMenuItem, len(a))
	for i, it := range a {
		out[i] = copyItem(it)
	}

	for _, itB := range b {
		if i, found := index[itB.Name]; found {
			out[i] = mergeItem(out[i], itB)
		} else {
			out = append(out, copyItem(itB))
		}
	}
	return out
}

// copyItem makes deep copy DataMenuItem
func copyItem(src DataMenuItem) DataMenuItem {
	dst := DataMenuItem{
		Name:     src.Name,
		Href:     src.Href,
		Attr:     src.Attr,
		Children: make([]DataMenuItem, len(src.Children)),
	}
	for i, ch := range src.Children {
		dst.Children[i] = copyItem(ch)
	}
	return dst
}
