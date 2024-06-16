package errorg

// JsonLog sets app's log option
func JsonLog() Options {
	handler := func() {
		defaultConfigs[log] = logTypeMap[jsonType]
	}
	return &options{handler}
}
