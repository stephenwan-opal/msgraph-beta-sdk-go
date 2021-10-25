package graph
import (
    "strings"
    "errors"
)
type RecordCompletionReason int

const (
    OPERATIONCANCELED_RECORDCOMPLETIONREASON RecordCompletionReason = iota
    STOPTONEDETECTED_RECORDCOMPLETIONREASON
    MAXRECORDDURATIONREACHED_RECORDCOMPLETIONREASON
    INITIALSILENCETIMEOUT_RECORDCOMPLETIONREASON
    MAXSILENCETIMEOUT_RECORDCOMPLETIONREASON
    PLAYPROMPTFAILED_RECORDCOMPLETIONREASON
    PLAYBEEPFAILED_RECORDCOMPLETIONREASON
    MEDIARECEIVETIMEOUT_RECORDCOMPLETIONREASON
    UNSPECIFIEDERROR_RECORDCOMPLETIONREASON
)

func (i RecordCompletionReason) String() string {
    return []string{"OPERATIONCANCELED", "STOPTONEDETECTED", "MAXRECORDDURATIONREACHED", "INITIALSILENCETIMEOUT", "MAXSILENCETIMEOUT", "PLAYPROMPTFAILED", "PLAYBEEPFAILED", "MEDIARECEIVETIMEOUT", "UNSPECIFIEDERROR"}[i]
}
func ParseRecordCompletionReason(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "OPERATIONCANCELED":
            return OPERATIONCANCELED_RECORDCOMPLETIONREASON, nil
        case "STOPTONEDETECTED":
            return STOPTONEDETECTED_RECORDCOMPLETIONREASON, nil
        case "MAXRECORDDURATIONREACHED":
            return MAXRECORDDURATIONREACHED_RECORDCOMPLETIONREASON, nil
        case "INITIALSILENCETIMEOUT":
            return INITIALSILENCETIMEOUT_RECORDCOMPLETIONREASON, nil
        case "MAXSILENCETIMEOUT":
            return MAXSILENCETIMEOUT_RECORDCOMPLETIONREASON, nil
        case "PLAYPROMPTFAILED":
            return PLAYPROMPTFAILED_RECORDCOMPLETIONREASON, nil
        case "PLAYBEEPFAILED":
            return PLAYBEEPFAILED_RECORDCOMPLETIONREASON, nil
        case "MEDIARECEIVETIMEOUT":
            return MEDIARECEIVETIMEOUT_RECORDCOMPLETIONREASON, nil
        case "UNSPECIFIEDERROR":
            return UNSPECIFIEDERROR_RECORDCOMPLETIONREASON, nil
    }
    return 0, errors.New("Unknown RecordCompletionReason value: " + v)
}
func SerializeRecordCompletionReason(values []RecordCompletionReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}