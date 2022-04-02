package funcs

import (
	"gopkg.in/yaml.v2"
)

func ToYaml(v interface{}) string {
	b, err := yaml.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func MustToYaml(v interface{}) (string, error) {
	b, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
