package funcs

import (
	"text/template"
)

func AttachToMap(baseMap map[string]any) map[string]any {
	baseMap["table"] = Table
	baseMap["toYaml"] = ToYaml
	baseMap["mustToYaml"] = MustToYaml
	return baseMap
}

func ClosureMap(baseMap map[string]any, t *template.Template) map[string]any {
	funcMap := AttachToMap(baseMap)
	funcMap["include"] = MakeInclude(t)
	funcMap["tpl"] = MakeTpl(t)
	return funcMap
}
