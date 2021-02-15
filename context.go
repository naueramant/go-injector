package injector

import (
	"reflect"
)

const (
	tagName = "inject"
)

type Context struct {
	typesProvided map[string]interface{}
	namedProvided map[string]interface{}
}

func New() *Context {
	return &Context{
		typesProvided: make(map[string]interface{}),
		namedProvided: make(map[string]interface{}),
	}
}

func (c *Context) Provide(val interface{}, names ...string) {
	if len(names) > 0 {
		for _, n := range names {
			c.namedProvided[n] = val
		}

		return
	}

	typeStr := reflect.TypeOf(val).String()
	c.typesProvided[typeStr] = val
}

func (c *Context) Remove(val interface{}) {
	var keys []string

	switch v := val.(type) {
	case string:
		keys = []string{v}
	case []string:
		keys = v
	default:
		keys = []string{
			reflect.TypeOf(val).String(),
		}
	}

	for _, k := range keys {
		delete(c.typesProvided, k)
		delete(c.namedProvided, k)
	}
}

func (c *Context) Inject(structPtr interface{}) error {
	targetType := reflect.TypeOf(structPtr)

	if !isStructPtr(structPtr) {
		return &ErrNonStructPointer{
			TargetType: targetType,
		}
	}

	chasedTargetType := targetType.Elem()

	for i := 0; i < chasedTargetType.NumField(); i++ {
		f := chasedTargetType.Field(i)

		var name string
		var m map[string]interface{}

		tagName, ok := f.Tag.Lookup(tagName)
		if ok {
			m = c.namedProvided
			name = tagName
		} else {
			m = c.typesProvided
			name = f.Type.String()
		}

		d, ok := m[name]
		if !ok {
			continue
		}

		targetField := reflect.ValueOf(structPtr).Elem().Field(i)
		injectValue := reflect.ValueOf(d)

		if targetField.Type() != injectValue.Type() {
			return &ErrFieldValueTypeMismatch{
				FieldName:  f.Name,
				TargetType: targetType,
				FieldType:  targetField.Type(),
				ValueType:  injectValue.Type(),
			}
		}

		targetField.Set(injectValue)
	}

	return nil
}

func (c *Context) Get(val interface{}) (res interface{}, ok bool) {
	switch v := val.(type) {
	case string:
		res, ok = c.namedProvided[v]
	default:
		res, ok = c.typesProvided[reflect.TypeOf(val).String()]
	}

	return
}

func (c *Context) Clone() (ctx *Context) {
	cp := *c

	return &cp
}
