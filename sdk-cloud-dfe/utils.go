package sdk_cloud_dfe

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"
)

func Encode(data string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	return encoded

}

func Decode(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(decoded)
	r, err := gzip.NewReader(buf)
	if err != nil {

		return string(decoded), nil
	}
	defer r.Close()

	decompressed, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(decompressed), nil
}

func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	return string(content), nil
}
