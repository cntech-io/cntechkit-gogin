package errorcodes

import (
    "encoding/json"
    "fmt"
)

type ErrorCodeList map[string]interface{}

var List ErrorCodeList = ErrorCodeList{
    ERRC10026.Code: ERRC10026,
    ERRC10123.Code : ERRC10123,
    ERRC10803.Code : ERRC10803,
    ERRC11646.Code: ERRC11646,
    ERRC14778.Code: ERRC14778,
    ERRC17006.Code: ERRC17006,
    ERRC17506.Code: ERRC17506,
    ERRC17666.Code: ERRC17666,
    ERRC19444.Code: ERRC19444,
    ERRC11677.Code: ERRC11677,
}

func (ecl *ErrorCodeList) ToJson() string {
    errorBytes, err := json.Marshal(ecl)
    if err != nil {
        fmt.Println("ErrorCodeList to Json conversion error!")
    }
    return string(errorBytes)
}