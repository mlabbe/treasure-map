package textfunc

import "encoding/xml"

func ToXml(v any) string {
	b, err := xml.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func MustToXml(v any) (string, error) {
	b, err := xml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func FromXml(s string) any {
	var v any
	err := xml.Unmarshal([]byte(s), &v)
	if err != nil {
		return nil
	}
	return v
}
