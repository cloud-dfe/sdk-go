package sdk_cloud_dfe

import (
	"errors"
	"regexp"
)

type base struct {
	Client client
}

func checkKey(payload map[string]interface{}) (string, error) {

	key, ok := payload["chave"].(string)

	if !ok {
		return "", errors.New("a chave deve ser uma string")
	}

	re := regexp.MustCompile(`[^0-9]`)
	key = re.ReplaceAllString(key, "")

	if len(key) != 44 {
		return "", errors.New("a chave deve ter 44 dígitos numéricos")
	}

	return key, nil
}
