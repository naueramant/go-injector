package injector

import (
	"fmt"
	"reflect"
)

type ErrNonStructPointer struct {
	TargetType reflect.Type
}

func (e *ErrNonStructPointer) Error() string {
	return fmt.Sprintf(
		"Can't inject into non-struct pointer of type '%s'",
		e.TargetType.String(),
	)
}

type ErrFieldValueTypeMismatch struct {
	TargetType reflect.Type
	FieldName  string
	FieldType  reflect.Type
	ValueType  reflect.Type
}

func (e *ErrFieldValueTypeMismatch) Error() string {
	return fmt.Sprintf(
		"Failed to inject into struct of type '%s' since field '%s' of type '%s' does not match injected value of type '%s'",
		e.TargetType.String(),
		e.FieldName,
		e.FieldType.String(),
		e.ValueType.String(),
	)
}
