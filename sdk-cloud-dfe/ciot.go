package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type ciot struct {
	Base base
}

func Ciot(b base) ciot {

	result := ciot{Base: b}

	return result
}

func (c ciot) Cria(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot", payload)

	return resp, err
}

func (c ciot) Consulta(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/ciot/%s", key), nil)

	return resp, err
}

func (c ciot) Busca(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/busca", payload)

	return resp, err
}

func (c ciot) Cancela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/cancela", payload)

	return resp, err
}

func (c ciot) Encerra(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/encerra", payload)

	return resp, err
}

func (c ciot) Retifica(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/retifica", payload)

	return resp, err
}

func (c ciot) Pdf(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/ciot/pdf/%s", key), nil)

	return resp, err
}

func (c ciot) Transportador(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/transportador", payload)

	return resp, err
}
