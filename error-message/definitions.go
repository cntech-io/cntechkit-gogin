package errormessage

import (
	"encoding/json"

	gogin "github.com/cntech-io/cntechkit-gogin/v2"
)

type Language struct {
	TR string `json:"tr"`
}

type ErrorMessage struct {
	Code     string   `json:"code"`
	Default  string   `json:"default"`
	Language Language `json:"languages"`
}

type ErrorMessageList map[string]interface{}

func (eml *ErrorMessageList) ToJson() string {
	errorBytes, err := json.Marshal(eml)
	if err != nil {
		gogin.LOGGER.Warn("ErrorCodeList to Json conversion error!")
	}
	return string(errorBytes)
}

func New(code string, message string) ErrorMessage {
	return ErrorMessage{
		Code:    code,
		Default: message,
	}
}
