package errorg

// TextLog sets app's log option
func TextLog() Options {
	handler := func() {
		defaultConfigs[log] = logTypeMap[textType]
	}
	return &options{handler}
}
