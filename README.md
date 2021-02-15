# go-injector

A lightweight context based go dependency injector.

# Example

**Struct injection**

In this example values are injected "by type" in to the targeted struct pointer.

_example:_

```go
package injector

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

**Named injections**

Instead of injection "by type" it is also possible to inject by tag name.

_example:_

```go
package injector

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

It is important that the provided type and field type match else an error will be returned by inject.
