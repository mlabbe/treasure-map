package funcs_test

import (
	"testing"
	"text/template"

	"github.com/bluebrown/treasure-map/funcs"
)

func TestMakeFuncMap(t *testing.T) {
	tpl := template.New("test")
	baseMap := make(map[string]any)
	fm := funcs.NewMap(baseMap, tpl)
	if _, ok := fm["include"]; !ok {
		t.Error("include is not found")
	}
}
