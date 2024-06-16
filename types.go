package errorg

const (
	errorgDirectory = ".errorg"
)

type optType int

const (
	log optType = iota + 1
)

// logType represents requested logType from user
type logType int

const (
	textType logType = iota + 1
	jsonType
)

var logTypeMap = map[logType]string{
	textType: "text",
	jsonType: "json",
}

type errorInfo struct {
	Msg string `json:"msg,omitempty"`
	Err error  `json:"err,omitempty"`
}

// GeneralConfig represents parsed yaml values
type GeneralConfig struct {
	Version     int       `yaml:"version"`
	Environment string    `yaml:"environment"`
	Plugins     []Plugins `yaml:"plugins"`
}
type Plugins struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
}
