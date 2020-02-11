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

func JSON(question string, stdout bool) (string, error) {
	value, err := String(question, stdout)
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
		err := json.Unmarshal([]byte(value), &obj)
		if err != nil {
			return "", fmt.Errorf("error unmarshaling array: %w", err)
		}
		compressed, err := json.Marshal(obj)
		if err != nil {
			return "", fmt.Errorf("error re-marshaling array: %w", err)
		}
		return string(compressed), nil
	case '{':
		obj := map[string]interface{}{}
		err := json.Unmarshal([]byte(value), &obj)
		if err != nil {
			return "", fmt.Errorf("error unmarshaling map: %w", err)
		}
		compressed, err := json.Marshal(obj)
		if err != nil {
			return "", fmt.Errorf("error re-marshaling map: %w", err)
		}
		return string(compressed), nil
	case '"':
		obj := ""
		err := json.Unmarshal([]byte(value), &obj)
		if err != nil {
			return "", fmt.Errorf("error unmarshaling string: %w", err)
		}
		return value, nil
	}

	obj := 0.0
	err = json.Unmarshal([]byte(value), &obj)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling float: %w", err)
	}
	return value, nil
}
