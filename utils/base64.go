package utils

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"
)

func EncodeBase64Image(imageBase64 string) (io.Reader, error) {
	coI := strings.Index(imageBase64, ",")
	rawImage := string(imageBase64)[coI+1:]

	unbased, err := base64.StdEncoding.DecodeString(string(rawImage))

	if err != nil {
		return nil, err
	}
	res := bytes.NewReader(unbased)

	return res, nil
}
