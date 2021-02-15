# go-injector

A lightweight context based go dependency injector.

## Examples

**Struct injection**

In this example values are injected "by type" in to the targeted struct pointer.

_example:_

```go
package main

import (
    "fmt"

    "github.com/naueramant/go-injector"
)

type Foo struct {
    A Bar
    B *Bar
    C string
    D int
}

type Bar struct{}

func main() {
    ctx := injector.New()

    bar := Bar{}

    ctx.Provide(bar)
    ctx.Provide(&bar)
    ctx.Provide("Hello World")
    ctx.Provide(42)

    foo := Foo{}

    if err := ctx.Inject(&foo); err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", foo)
}

```

_output:_

```sh
{A:{} B:0x58b470 C:Hello World D:42}
```

Test it yourself at the [go playground](https://play.golang.org/p/EnKDlfwPp5A).

**Named injections**

Instead of injection "by type" it is also possible to inject by tag name.

_example:_

```go
package main

import (
    "fmt"

    "github.com/naueramant/go-injector"
)

type Foo struct {
    A string `inject:"foo"`
    B string
}

func main() {
    ctx := injector.New()

    ctx.Provide("Hello World", "foo")
    ctx.Provide("bar")

    foo := Foo{}

    if err := ctx.Inject(&foo); err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", foo)
}

```

_output:_

```sh
{A:Hello World B:bar}
```

Test it yourself at the [go playground](https://play.golang.org/p/vZj6jGufmfQ).

It is important that the provided value type and field type match else an error will be returned by inject.

**Global context**

A global context is provided for convenience and can be accessed directly on the package as so:

```go
package main

import (
    "fmt"

    "github.com/naueramant/go-injector"
)

type Foo struct {
    A int
}

func main() {
    injector.Provide(42)

    foo := Foo{}

    if err := injector.Inject(&foo); err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", foo)
}
```

```sh
{A:42}
```

Test it yourself at the [go playground](https://play.golang.org/p/odoAHwMW1Tw).

**Struct tags**

```go
type Foo struct {
    A string `inject:"foo"`          // named
    B string `inject:"foo,required"` // named and required
    C string `inject:"required"`     // required
    D string `inject:"-"`            // skip
}
```

If a field is required and not provided `inject` will return an error.
