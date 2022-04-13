# Treasure Map

A collection of go template functions. Some functions are inspired by [helm](https://helm.sh/), i.e. include and tpl.

## Usage

Some function require to close over the template itself. Therefore, the MakeFuncMap function takes the template as argument and returns the fully constructed map. In the below example [sprigs](http://masterminds.github.io/sprig/) TxtFuncMap is passed as base map.

```go
import (
    "testing"
    "text/template"

    "github.com/Masterminds/sprig/v3"
    "github.com/bluebrown/treasure-map/funcs"
)

func TestMakeFuncMap(t *testing.T) {
    tpl := template.New(rootTemplateName)
    tpl.Funcs(funcs.ClosureMap(sprig.TxtFuncMap(), tpl))
}
```

## Functions

Function    | Description                                         | Returns       | Example
------------|-----------------------------------------------------|---------------|------------------------------
include     | render a previously associated template             | string        | `{{ include "my-helper" . }}`
tpl         | render a template snippet                           | string        | `{{ tpl "{{ .foo.bar }}" . }}`
toYaml      | convert to yaml                                     | string        | `{{ toYaml . }}`
mustToYaml  | convert to yaml, errors if encoding fails           | string, error | `{{ mustToYaml . }}`
table       | convert to table, list of object or list of lists   | string        | `{{ table . }}`
