package main

type stdGoError struct {
}

func (g stdGoError) GetErrorMap() map[string]interface{} {
	return map[string]interface{}{
		"100": struct {
			Message string `json:"msg"`
		}{"test"},
	}
}

var Group stdGoError
