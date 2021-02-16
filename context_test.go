package injector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	ctx := New()

	assert.NotNil(t, ctx.namedProvided)
	assert.NotNil(t, ctx.typesProvided)

	assert.Empty(t, ctx.typesProvided)
	assert.Empty(t, ctx.namedProvided)
}

func TestContext_Provide_Bool(t *testing.T) {
	t.Parallel()

	ctx := New()
	ctx.Provide(bool(true))
	assert.NotNil(t, ctx.typesProvided["bool"])
}

func TestContext_Provide_String(t *testing.T) {
	t.Parallel()

	ctx := New()
	ctx.Provide(string("Hello world"))
	assert.NotNil(t, ctx.typesProvided["string"])
}

func TestContext_Provide_Int(t *testing.T) {
	t.Parallel()

	ctx := New()
	ctx.Provide(int(1))
	assert.NotNil(t, ctx.typesProvided["int"])

	ctx = New()
	ctx.Provide(int8(1))
	assert.NotNil(t, ctx.typesProvided["int8"])

	ctx = New()
	ctx.Provide(int16(1))
	assert.NotNil(t, ctx.typesProvided["int16"])

	ctx = New()
	ctx.Provide(int32(1))
	assert.NotNil(t, ctx.typesProvided["int32"])

	ctx = New()
	ctx.Provide(int64(1))
	assert.NotNil(t, ctx.typesProvided["int64"])
}

func TestContext_Provide_Uint(t *testing.T) {
	t.Parallel()

	ctx := New()
	ctx.Provide(uint(1))
	assert.NotNil(t, ctx.typesProvided["uint"])

	ctx = New()
	ctx.Provide(uint8(1))
	assert.NotNil(t, ctx.typesProvided["uint8"])

	ctx = New()
	ctx.Provide(uint16(1))
	assert.NotNil(t, ctx.typesProvided["uint16"])

	ctx = New()
	ctx.Provide(uint32(1))
	assert.NotNil(t, ctx.typesProvided["uint32"])

	ctx = New()
	ctx.Provide(uint64(1))
	assert.NotNil(t, ctx.typesProvided["uint64"])

	ctx = New()
	ctx.Provide(uintptr(1))
	assert.NotNil(t, ctx.typesProvided["uintptr"])
}

func TestContext_Provide_Float(t *testing.T) {
	t.Parallel()

	ctx := New()
	ctx.Provide(float32(1))
	assert.NotNil(t, ctx.typesProvided["float32"])

	ctx = New()
	ctx.Provide(float64(1))
	assert.NotNil(t, ctx.typesProvided["float64"])
}

func TestContext_Provide_Byte(t *testing.T) {
	t.Parallel()

	ctx := New()
	ctx.Provide(byte(1))
	assert.NotNil(t, ctx.typesProvided["uint8"])
}

func TestContext_Provide_Rune(t *testing.T) {
	t.Parallel()

	ctx := New()
	ctx.Provide(rune(1))
	assert.NotNil(t, ctx.typesProvided["int32"])
}

func TestContext_Provide_Complex(t *testing.T) {
	t.Parallel()

	ctx := New()
	ctx.Provide(complex64(1))
	assert.NotNil(t, ctx.typesProvided["complex64"])

	ctx = New()
	ctx.Provide(complex128(1))
	assert.NotNil(t, ctx.typesProvided["complex128"])
}

func TestContext_Provide_Struct(t *testing.T) {
	t.Parallel()

}
