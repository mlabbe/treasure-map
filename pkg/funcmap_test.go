package treasuremap_test

import (
	"testing"
	"text/template"

	treasuremap "github.com/bluebrown/treasure-map/pkg"
)

func TestMakeFuncMap(t *testing.T) {
	tpl := template.New("test")
	baseMap := make(map[string]any)
	fm := treasuremap.MakeFuncMap(baseMap, tpl)
	if _, ok := fm["include"]; !ok {
		t.Error("include is not found")
	}
}
