package injector

var globalCtx = New()

func Provide(val interface{}, names ...string) {
	globalCtx.Provide(val, names...)
}

func Remove(val interface{}) {
	globalCtx.Remove(val)
}

func Inject(structPtr interface{}) error {
	return globalCtx.Inject(structPtr)
}

func Get(val interface{}) (res interface{}, ok bool) {
	return globalCtx.Get(val)
}

func Clone() (ctx *Context) {
	return globalCtx.Clone()
}
