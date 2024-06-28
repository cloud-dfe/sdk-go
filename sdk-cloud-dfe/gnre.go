package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type gnre struct {
	Base base
}

func Gnre(b base) gnre {

	result := gnre{Base: b}

	return result
}

func (c gnre) Consulta(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/gnre/%s", key), nil)

	return resp, err
}

func (c gnre) Cria(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/gnre", payload)

	return resp, err
}

func (c gnre) ConfigUf(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/gnre/configuf", payload)

	return resp, err
}
