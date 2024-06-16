package errorg

// defaultConfigs holds app's default config
// ex: if any config is not set while app initializing text type will be default
var defaultConfigs = map[optType]string{
	log: logTypeMap[textType],
}

type options struct {
	handlerFunc func()
}

func (o options) Exec() {
	o.handlerFunc()
}

// Options represents options that needs to be implemented by whole options
type Options interface {
	Exec()
}

// optionsInitializer executes the options
func optionsInitializer(options []Options) {
	for idx := range options {
		options[idx].Exec()
	}
}
