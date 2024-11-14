package generator

type OutputMessage struct {
	Name  string
	HasId bool
}

type Output struct {
	HeaderComments []string
	PackageName    string
	ProtoPackage   string
	Imports        []string
	Messages       []OutputMessage
}

type OutputPropertySetter interface {
	apply(*Output)
}

type OutputPropertySetterFunc func(*Output)

func (o OutputPropertySetterFunc) apply(output *Output) {
	o(output)
}

func WithPackageName(packageName string) OutputPropertySetterFunc {
	return func(o *Output) {
		o.PackageName = packageName
	}
}

func WithProtoPackage(protoPackage string) OutputPropertySetterFunc {
	return func(o *Output) {
		o.ProtoPackage = protoPackage
	}
}

func WithHeaderComments(comments []string) OutputPropertySetterFunc {
	return func(o *Output) {
		o.HeaderComments = comments
	}
}

func WithImports(imports []string) OutputPropertySetterFunc {
	return func(o *Output) {
		o.Imports = imports
	}
}

func WithMessages(messages []OutputMessage) OutputPropertySetterFunc {
	return func(o *Output) {
		o.Messages = messages
	}
}

func AppendMessage(message OutputMessage) OutputPropertySetterFunc {
	return func(o *Output) {
		o.Messages = append(o.Messages, message)
	}
}

func NewOutput(opts ...OutputPropertySetter) *Output {
	output := &Output{}

	for _, opt := range opts {
		opt.apply(output)
	}

	return output
}
