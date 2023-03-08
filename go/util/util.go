package util

import "encoding/json"

func Jsonify(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
