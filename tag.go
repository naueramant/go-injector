package injector

import (
	"reflect"
	"strings"
)

const (
	tagIdentifier = "inject"
)

type Tag struct {
	Name     string
	Required bool
	Skip     bool
}

func parseTag(st reflect.StructTag) (*Tag, bool) {
	val, ok := st.Lookup("inject")
	if !ok {
		return nil, ok
	}

	val = strings.TrimSpace(val)

	vals := strings.Split(val, ",")
	if len(vals) == 0 {
		return nil, false
	}

	var name string
	if vals[0] != "-" && vals[0] != "required" {
		name = vals[0]
	}

	return &Tag{
		Name:     name,
		Required: strings.Contains(val, "required"),
		Skip:     strings.Contains(val, "-"),
	}, true
}
