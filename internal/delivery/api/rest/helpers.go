package rest

import (
	"net/url"
	"strconv"
)

func parseInt(values url.Values, key string) (int, error) {
	value := values.Get(key)

	valueI := int(0)

	if value != "" {
		_valueI, err := strconv.ParseUint(value, 10, 32)
		if err != nil {

		}
		valueI = int(_valueI)
	}

	return valueI, nil
}
