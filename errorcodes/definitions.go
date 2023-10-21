package errorcodes

import (
    "encoding/json"
    "fmt"
)

type Language struct {
    TR string `json:"tr"`
}


type ErrorCode struct {
    Code string `json:"code"`
    Default string `json:"default"`
    Language Language `json:"languages"`
}


func (ec *ErrorCode) ToJson() string {
    errorBytes, err := json.Marshal(ec)
    if err != nil {
        fmt.Println("ErrorCode to Json conversion error!")
    }
    return string(errorBytes)
}
