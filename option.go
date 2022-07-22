package copier

type Option struct {
	// setting this value to true will ignore copying zero values of all the fields, including bools, as well as a
	// struct having all it's fields set to their zero values respectively (see IsZero() in reflect/value.go)
	IgnoreEmpty    bool
	DeepCopy       bool
	Converters     []TypeConverter
	InitNilSlice   bool
	InitNilPointer bool
}

type TypeConverter struct {
	SrcType interface{}
	DstType interface{}
	Fn      func(src interface{}) (interface{}, error)
}

type Optional func(o *Option)

func WithIgnoreEmptyOption() Optional {
	return func(o *Option) {
		o.IgnoreEmpty = true
	}
}

func WithDeepCopyOption() Optional {
	return func(o *Option) {
		o.DeepCopy = true
	}
}

func WithConverters(in []TypeConverter) Optional {
	return func(o *Option) {
		o.Converters = in
	}
}

func WithInitNilSlice() Optional {
	return func(o *Option) {
		o.InitNilSlice = true
	}
}

func WithInitNilPointer() Optional {
	return func(o *Option) {
		o.InitNilPointer = true
	}
}

func defaultOption() *Option {
	return &Option{}
}
