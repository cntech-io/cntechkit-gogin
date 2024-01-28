package errormessage

import (
	"encoding/json"
	"fmt"
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

func (em *ErrorMessage) ToJson() string {
	errorBytes, err := json.Marshal(em)
	if err != nil {
		fmt.Println("ErrorCode to Json conversion error!")
	}
	return string(errorBytes)
}

func (eml *ErrorMessageList) ToJson() string {
	errorBytes, err := json.Marshal(eml)
	if err != nil {
		fmt.Println("ErrorCodeList to Json conversion error!")
	}
	return string(errorBytes)
}

func New(code string, message string) ErrorMessage {
	return ErrorMessage{
		Code:    code,
		Default: message,
	}
}
