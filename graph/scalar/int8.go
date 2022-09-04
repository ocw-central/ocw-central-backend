package graph

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"

	"github.com/kafugen/ocwcentral/utils"
)

func MarshalInt8(i int8) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.FormatInt(int64(i), 10))

		// TODO: use logger to record the error
		// ignore error to avoid breaking the response and fix lint error
		utils.IgnoreErr(err)
	})
}

func UnmarshalInt8(v interface{}) (int8, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return 0, err
		}
		return int8(iv), nil
	case int:
		return int8(v), nil
	case int64:
		return int8(v), nil
	case json.Number:
		iv, err := strconv.ParseInt(string(v), 10, 8)
		if err != nil {
			return 0, err
		}
		return int8(iv), nil
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
