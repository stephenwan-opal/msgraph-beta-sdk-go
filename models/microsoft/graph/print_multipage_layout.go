package graph
import (
    "strings"
    "errors"
)
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
    return []string{"CLOCKWISEFROMTOPLEFT", "COUNTERCLOCKWISEFROMTOPLEFT", "COUNTERCLOCKWISEFROMTOPRIGHT", "CLOCKWISEFROMTOPRIGHT", "COUNTERCLOCKWISEFROMBOTTOMLEFT", "CLOCKWISEFROMBOTTOMLEFT", "COUNTERCLOCKWISEFROMBOTTOMRIGHT", "CLOCKWISEFROMBOTTOMRIGHT"}[i]
}
func ParsePrintMultipageLayout(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CLOCKWISEFROMTOPLEFT":
            return CLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT, nil
        case "COUNTERCLOCKWISEFROMTOPLEFT":
            return COUNTERCLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT, nil
        case "COUNTERCLOCKWISEFROMTOPRIGHT":
            return COUNTERCLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT, nil
        case "CLOCKWISEFROMTOPRIGHT":
            return CLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT, nil
        case "COUNTERCLOCKWISEFROMBOTTOMLEFT":
            return COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT, nil
        case "CLOCKWISEFROMBOTTOMLEFT":
            return CLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT, nil
        case "COUNTERCLOCKWISEFROMBOTTOMRIGHT":
            return COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT, nil
        case "CLOCKWISEFROMBOTTOMRIGHT":
            return CLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT, nil
    }
    return 0, errors.New("Unknown PrintMultipageLayout value: " + v)
}
func SerializePrintMultipageLayout(values []PrintMultipageLayout) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}