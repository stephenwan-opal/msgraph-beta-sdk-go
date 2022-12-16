package models
import (
    "errors"
)
// Provides operations to manage the collection of user entities.
type RatingJapanTelevisionType int

const (
    // Default value, allow all TV shows content
    ALLALLOWED_RATINGJAPANTELEVISIONTYPE RatingJapanTelevisionType = iota
    // Do not allow any TV shows content
    ALLBLOCKED_RATINGJAPANTELEVISIONTYPE
    // All TV content is explicitly allowed
    EXPLICITALLOWED_RATINGJAPANTELEVISIONTYPE
)

func (i RatingJapanTelevisionType) String() string {
    return []string{"allAllowed", "allBlocked", "explicitAllowed"}[i]
}
func ParseRatingJapanTelevisionType(v string) (interface{}, error) {
    result := ALLALLOWED_RATINGJAPANTELEVISIONTYPE
    switch v {
        case "allAllowed":
            result = ALLALLOWED_RATINGJAPANTELEVISIONTYPE
        case "allBlocked":
            result = ALLBLOCKED_RATINGJAPANTELEVISIONTYPE
        case "explicitAllowed":
            result = EXPLICITALLOWED_RATINGJAPANTELEVISIONTYPE
        default:
            return 0, errors.New("Unknown RatingJapanTelevisionType value: " + v)
    }
    return &result, nil
}
func SerializeRatingJapanTelevisionType(values []RatingJapanTelevisionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
