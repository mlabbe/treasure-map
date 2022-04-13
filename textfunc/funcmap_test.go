package textfunc_test

import (
	"testing"
	"text/template"

	"github.com/bluebrown/treasure-map/textfunc"
)

func TestMapClosure(t *testing.T) {
	tpl := template.New("test")
	baseMap := make(map[string]any)
	fm := textfunc.MapClosure(baseMap, tpl)
	if _, ok := fm["include"]; !ok {
		t.Error("include is not found")
	}
}

func TestTextFuncMap(t *testing.T) {
	fm := textfunc.Map()
	if _, ok := fm["toYaml"]; !ok {
		t.Error("include is not found")
	}
}
