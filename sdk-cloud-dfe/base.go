package sdk_cloud_dfe

import (
	"regexp"
)

type base struct {
	Client client
}

func checkKey(payload map[string]string) (string, error) {

	key := payload["chave"]

	re := regexp.MustCompile(`[^0-9]`)
	key = re.ReplaceAllString(key, "")

	if len(key) != 44 {
		panic("A chave deve ter 44 dígitos numéricos.")
	}

	return key, nil
}
