package utils

import (
	"encoding/base64"
	"errors"
)

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", errors.New("can't decode base64")
	}
	return string(data), nil
}
