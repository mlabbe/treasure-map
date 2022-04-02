package treasuremap

import (
	"text/template"

	"github.com/bluebrown/treasure-map/pkg/funcs"
)

func AttachFuncs(baseMap map[string]any) map[string]any {
	baseMap["table"] = funcs.Table
	baseMap["toYaml"] = funcs.ToYaml
	baseMap["mustToYaml"] = funcs.MustToYaml
	return baseMap
}

func MakeFuncMap(baseMap map[string]any, t *template.Template) map[string]any {
	funcMap := AttachFuncs(baseMap)
	funcMap["include"] = funcs.MakeInclude(t)
	funcMap["tpl"] = funcs.MakeTpl(t)
	return funcMap
}
