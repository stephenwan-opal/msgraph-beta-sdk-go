package models
import (
    "strings"
    "errors"
)
// Provides operations to call the redirect method.
type CallDisposition int

const (
    DEFAULT_ESCAPED_CALLDISPOSITION CallDisposition = iota
    SIMULTANEOUSRING_CALLDISPOSITION
    FORWARD_CALLDISPOSITION
)

func (i CallDisposition) String() string {
    return []string{"DEFAULT_ESCAPED", "SIMULTANEOUSRING", "FORWARD"}[i]
}
func ParseCallDisposition(v string) (interface{}, error) {
    result := DEFAULT_ESCAPED_CALLDISPOSITION
    switch strings.ToUpper(v) {
        case "DEFAULT_ESCAPED":
            result = DEFAULT_ESCAPED_CALLDISPOSITION
        case "SIMULTANEOUSRING":
            result = SIMULTANEOUSRING_CALLDISPOSITION
        case "FORWARD":
            result = FORWARD_CALLDISPOSITION
        default:
            return 0, errors.New("Unknown CallDisposition value: " + v)
    }
    return &result, nil
}
func SerializeCallDisposition(values []CallDisposition) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}