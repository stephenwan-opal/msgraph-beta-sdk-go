package graph
import (
    "strings"
    "errors"
)
type CompanyPortalAction int

const (
    UNKNOWN_COMPANYPORTALACTION CompanyPortalAction = iota
    REMOVE_COMPANYPORTALACTION
    RESET_COMPANYPORTALACTION
)

func (i CompanyPortalAction) String() string {
    return []string{"UNKNOWN", "REMOVE", "RESET"}[i]
}
func ParseCompanyPortalAction(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_COMPANYPORTALACTION, nil
        case "REMOVE":
            return REMOVE_COMPANYPORTALACTION, nil
        case "RESET":
            return RESET_COMPANYPORTALACTION, nil
    }
    return 0, errors.New("Unknown CompanyPortalAction value: " + v)
}
func SerializeCompanyPortalAction(values []CompanyPortalAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}