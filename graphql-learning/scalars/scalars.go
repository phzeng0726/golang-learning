package scalars

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalDate(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(time.DateOnly)))
	})
}

func UnmarshalDate(v interface{}) (time.Time, error) {
	value, ok := v.(string)

	if !ok {
		return time.Time{}, errors.New("error parsing time")
	}

	t, err := time.Parse(time.DateOnly, value)

	if err != nil {
		return t, err
	}
	return t.UTC(), nil

}
