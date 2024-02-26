package query

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ToGraphData used to convert any value to a graph-compatible representation
func ToGraphData(i interface{}) (string, error) {
	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.String:
		return strconv.Quote(v.String()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	case reflect.Slice, reflect.Array:
		var stringArr []string

		for i := 0; i < v.Len(); i++ {
			eleStr, err := ToGraphData(v.Index(i).Interface())
			if err != nil {
				return "", err
			}

			stringArr = append(stringArr, eleStr)
		}

		return "[" + strings.Join(stringArr, ",") + "]", nil
	default:
		return "", fmt.Errorf("unsupported kind %s", v.Kind())
	}
}
