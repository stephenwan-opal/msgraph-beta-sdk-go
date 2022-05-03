package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the compliance singleton.
type PrintPresentationDirection int

const (
    CLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION PrintPresentationDirection = iota
    COUNTERCLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION
    COUNTERCLOCKWISEFROMTOPRIGHT_PRINTPRESENTATIONDIRECTION
    CLOCKWISEFROMTOPRIGHT_PRINTPRESENTATIONDIRECTION
    COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTPRESENTATIONDIRECTION
    CLOCKWISEFROMBOTTOMLEFT_PRINTPRESENTATIONDIRECTION
    COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTPRESENTATIONDIRECTION
    CLOCKWISEFROMBOTTOMRIGHT_PRINTPRESENTATIONDIRECTION
)

func (i PrintPresentationDirection) String() string {
    return []string{"CLOCKWISEFROMTOPLEFT", "COUNTERCLOCKWISEFROMTOPLEFT", "COUNTERCLOCKWISEFROMTOPRIGHT", "CLOCKWISEFROMTOPRIGHT", "COUNTERCLOCKWISEFROMBOTTOMLEFT", "CLOCKWISEFROMBOTTOMLEFT", "COUNTERCLOCKWISEFROMBOTTOMRIGHT", "CLOCKWISEFROMBOTTOMRIGHT"}[i]
}
func ParsePrintPresentationDirection(v string) (interface{}, error) {
    result := CLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION
    switch strings.ToUpper(v) {
        case "CLOCKWISEFROMTOPLEFT":
            result = CLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION
        case "COUNTERCLOCKWISEFROMTOPLEFT":
            result = COUNTERCLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION
        case "COUNTERCLOCKWISEFROMTOPRIGHT":
            result = COUNTERCLOCKWISEFROMTOPRIGHT_PRINTPRESENTATIONDIRECTION
        case "CLOCKWISEFROMTOPRIGHT":
            result = CLOCKWISEFROMTOPRIGHT_PRINTPRESENTATIONDIRECTION
        case "COUNTERCLOCKWISEFROMBOTTOMLEFT":
            result = COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTPRESENTATIONDIRECTION
        case "CLOCKWISEFROMBOTTOMLEFT":
            result = CLOCKWISEFROMBOTTOMLEFT_PRINTPRESENTATIONDIRECTION
        case "COUNTERCLOCKWISEFROMBOTTOMRIGHT":
            result = COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTPRESENTATIONDIRECTION
        case "CLOCKWISEFROMBOTTOMRIGHT":
            result = CLOCKWISEFROMBOTTOMRIGHT_PRINTPRESENTATIONDIRECTION
        default:
            return 0, errors.New("Unknown PrintPresentationDirection value: " + v)
    }
    return &result, nil
}
func SerializePrintPresentationDirection(values []PrintPresentationDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
