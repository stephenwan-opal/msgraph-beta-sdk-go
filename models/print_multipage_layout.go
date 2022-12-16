package models
import (
    "errors"
)
// Provides operations to manage the collection of user entities.
type PrintMultipageLayout int

const (
    CLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT PrintMultipageLayout = iota
    COUNTERCLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT
    COUNTERCLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT
    CLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT
    COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT
    CLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT
    COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT
    CLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT
)

func (i PrintMultipageLayout) String() string {
    return []string{"clockwiseFromTopLeft", "counterclockwiseFromTopLeft", "counterclockwiseFromTopRight", "clockwiseFromTopRight", "counterclockwiseFromBottomLeft", "clockwiseFromBottomLeft", "counterclockwiseFromBottomRight", "clockwiseFromBottomRight"}[i]
}
func ParsePrintMultipageLayout(v string) (interface{}, error) {
    result := CLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT
    switch v {
        case "clockwiseFromTopLeft":
            result = CLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT
        case "counterclockwiseFromTopLeft":
            result = COUNTERCLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT
        case "counterclockwiseFromTopRight":
            result = COUNTERCLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT
        case "clockwiseFromTopRight":
            result = CLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT
        case "counterclockwiseFromBottomLeft":
            result = COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT
        case "clockwiseFromBottomLeft":
            result = CLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT
        case "counterclockwiseFromBottomRight":
            result = COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT
        case "clockwiseFromBottomRight":
            result = CLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT
        default:
            return 0, errors.New("Unknown PrintMultipageLayout value: " + v)
    }
    return &result, nil
}
func SerializePrintMultipageLayout(values []PrintMultipageLayout) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
