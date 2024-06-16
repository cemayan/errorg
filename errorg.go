package errorg

import (
	"encoding/json"
)

// errorsMap holds errors that implemented group's errors
var errorsMap = map[string]map[string]interface{}{}

// errorg represents errorg struct
// Groyp represents the group ex: http,go etc..
// Code represents the specific code for given groyp ex:  "http","404" => "Not found"
type errorg struct {
	Group string
	Code  string
}

// Error return error according to given information's
func (e errorg) Error() string {
	return e.prepare()
}

// prepare returns result string that is changed
func (e errorg) prepare() string {

	var errInfo errorInfo

	currentLog := defaultConfigs[log]
	errStruct := errorsMap[e.Group][e.Code]
	marshal, _ := json.Marshal(errStruct)

	if currentLog == logTypeMap[jsonType] {
		return string(marshal)
	}

	err := json.Unmarshal(marshal, &errInfo)
	if err != nil {
		return "an error occurred  when unmarshalling error info: " + err.Error()
	}
	return errInfo.Msg
}

// Init initializes the errorg
// This method accepts options hereby it can be set options
func Init(options ...Options) {

	directoryInitializer()
	configInitializer()
	optionsInitializer(options)
	groupsInitializer()
}

// New return builtin error according to given values
func New(group string, code string) error {
	return &errorg{Group: group, Code: code}
}
