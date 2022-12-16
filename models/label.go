package models
import (
    "errors"
)
// Provides operations to manage the collection of user entities.
type Label int

const (
    TITLE_LABEL Label = iota
    URL_LABEL
    CREATEDBY_LABEL
    LASTMODIFIEDBY_LABEL
    AUTHORS_LABEL
    CREATEDDATETIME_LABEL
    LASTMODIFIEDDATETIME_LABEL
    FILENAME_LABEL
    FILEEXTENSION_LABEL
)

func (i Label) String() string {
    return []string{"title", "url", "createdBy", "lastModifiedBy", "authors", "createdDateTime", "lastModifiedDateTime", "fileName", "fileExtension"}[i]
}
func ParseLabel(v string) (interface{}, error) {
    result := TITLE_LABEL
    switch v {
        case "title":
            result = TITLE_LABEL
        case "url":
            result = URL_LABEL
        case "createdBy":
            result = CREATEDBY_LABEL
        case "lastModifiedBy":
            result = LASTMODIFIEDBY_LABEL
        case "authors":
            result = AUTHORS_LABEL
        case "createdDateTime":
            result = CREATEDDATETIME_LABEL
        case "lastModifiedDateTime":
            result = LASTMODIFIEDDATETIME_LABEL
        case "fileName":
            result = FILENAME_LABEL
        case "fileExtension":
            result = FILEEXTENSION_LABEL
        default:
            return 0, errors.New("Unknown Label value: " + v)
    }
    return &result, nil
}
func SerializeLabel(values []Label) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
