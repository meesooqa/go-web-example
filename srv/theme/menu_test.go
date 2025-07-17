package theme

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeMenu_Simple(t *testing.T) {
	mainMenu1 := map[string]DataMenuItem{
		"Main": {
			Name: "Main1",
			Href: "/main1",
			Attr: "data-test=\"m1\"",
			Children: []DataMenuItem{{
				Name: "Home",
				Href: "/",
				Attr: "title=\"Home\"",
			}},
		},
	}
	mainMenu2 := map[string]DataMenuItem{
		"Main": {
			Name: "Main2", // should be ignored since first has Name
			Children: []DataMenuItem{{
				Name: "Demo",
				Href: "/demo",
				Children: []DataMenuItem{{
					Name: "Sub Item 1",
					Href: "/demo/page1",
				}, {
					Name: "Sub Item 2",
					Href: "/demo/page2",
				}},
			}},
		},
	}
	merged := mergeMenu(mainMenu1, mainMenu2)

	// Assert root key exists
	assert.Contains(t, merged, "Main")
	item := merged["Main"]

	// Name, Href, Attr should come from first map
	assert.Equal(t, "Main1", item.Name)
	assert.Equal(t, "/main1", item.Href)
	assert.Equal(t, template.HTMLAttr("data-test=\"m1\""), item.Attr)

	// Children length and order
	assert.Len(t, item.Children, 2)
	assert.Equal(t, "Home", item.Children[0].Name)
	assert.Equal(t, "Demo", item.Children[1].Name)

	// Demo children
	demoChildren := item.Children[1].Children
	assert.Len(t, demoChildren, 2)
	assert.Equal(t, "Sub Item 1", demoChildren[0].Name)
	assert.Equal(t, "/demo/page1", demoChildren[0].Href)
	assert.Equal(t, "Sub Item 2", demoChildren[1].Name)
}

func TestMergeMenu_EmptyAndNil(t *testing.T) {
	// empty input
	merged := mergeMenu()
	assert.Empty(t, merged)

	// maps with empty children
	m1 := map[string]DataMenuItem{"A": {Name: "A"}}
	m2 := map[string]DataMenuItem{"A": {}}
	res := mergeMenu(m1, m2)
	assert.Len(t, res, 1)
	assert.Equal(t, "A", res["A"].Name)
}

func TestMergeMenu_MultipleKeys(t *testing.T) {
	m1 := map[string]DataMenuItem{"X": {Name: "X1"}}
	m2 := map[string]DataMenuItem{"Y": {Name: "Y1"}}
	res := mergeMenu(m1, m2)

	assert.Len(t, res, 2)
	assert.Equal(t, "X1", res["X"].Name)
	assert.Equal(t, "Y1", res["Y"].Name)
}

func TestMergeMenu_DeepMerge(t *testing.T) {
	// deeper children merging
	m1 := map[string]DataMenuItem{
		"Root": {Children: []DataMenuItem{{Name: "A", Children: []DataMenuItem{{Name: "A1"}}}}},
	}
	m2 := map[string]DataMenuItem{
		"Root": {Children: []DataMenuItem{{Name: "A", Children: []DataMenuItem{{Name: "A2"}}}}},
	}
	res := mergeMenu(m1, m2)
	children := res["Root"].Children
	assert.Len(t, children, 1)
	assert.Equal(t, "A", children[0].Name)
	// merged grandchildren
	gc := children[0].Children
	assert.Len(t, gc, 2)
	assert.ElementsMatch(t,
		[]string{gc[0].Name, gc[1].Name},
		[]string{"A1", "A2"},
	)
}

func TestSortMenu_Simple(t *testing.T) {
	items := []DataMenuItem{
		{Name: "Second", Sort: 2},
		{Name: "First", Sort: 1},
		{Name: "Third", Sort: 3},
	}
	sortMenu(items)

	assert.Equal(t, "First", items[0].Name)
	assert.Equal(t, "Second", items[1].Name)
	assert.Equal(t, "Third", items[2].Name)
}

func TestSortMenu_Nested(t *testing.T) {
	items := []DataMenuItem{
		{
			Name: "Parent",
			Sort: 1,
			Children: []DataMenuItem{
				{Name: "ChildB", Sort: 2},
				{Name: "ChildA", Sort: 1},
			},
		},
	}
	sortMenu(items)

	// Parent stays
	assert.Equal(t, "Parent", items[0].Name)
	// Children sorted
	children := items[0].Children
	assert.Equal(t, "ChildA", children[0].Name)
	assert.Equal(t, "ChildB", children[1].Name)
}
