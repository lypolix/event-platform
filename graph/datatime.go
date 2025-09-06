package graph

import (
    "time"
    "fmt"
    "github.com/99designs/gqlgen/graphql"
)

func MarshalDateTime(t time.Time) graphql.Marshaler {
    return graphql.MarshalString(t.Format(time.RFC3339))
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
    str, ok := v.(string)
    if !ok {
        return time.Time{}, fmt.Errorf("DateTime must be a string")
    }
    return time.Parse(time.RFC3339, str)
}
