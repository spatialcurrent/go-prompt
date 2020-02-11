// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package prompt

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

func SecretJSON(question string, stdout bool) (string, error) {
	value, err := SecretString(question, stdout)
	if err != nil {
		return "", err
	}
	if value == "true" || value == "false" || value == "null" {
		return value, nil
	}

	first, _ := utf8.DecodeRune([]byte(value))
	if first == utf8.RuneError {
		return "", fmt.Errorf("invalid string %q", value)
	}

	switch first {
	case '[':
		obj := make([]interface{}, 0)
		errUnmarshal := json.Unmarshal([]byte(value), &obj)
		if errUnmarshal != nil {
			return "", fmt.Errorf("error unmarshaling array: %w", errUnmarshal)
		}
		compressed, errMarshal := json.Marshal(obj)
		if errMarshal != nil {
			return "", fmt.Errorf("error re-marshaling array: %w", errMarshal)
		}
		return string(compressed), nil
	case '{':
		obj := map[string]interface{}{}
		errUnmarshal := json.Unmarshal([]byte(value), &obj)
		if errUnmarshal != nil {
			return "", fmt.Errorf("error unmarshaling map: %w", errUnmarshal)
		}
		compressed, errMarshal := json.Marshal(obj)
		if errMarshal != nil {
			return "", fmt.Errorf("error re-marshaling map: %w", errMarshal)
		}
		return string(compressed), nil
	case '"':
		obj := ""
		errUnmarshal := json.Unmarshal([]byte(value), &obj)
		if errUnmarshal != nil {
			return "", fmt.Errorf("error unmarshaling string: %w", errUnmarshal)
		}
		return value, nil
	}

	obj := 0.0
	errUnmarshal := json.Unmarshal([]byte(value), &obj)
	if errUnmarshal != nil {
		return "", fmt.Errorf("error unmarshaling float: %w", errUnmarshal)
	}
	return value, nil
}
