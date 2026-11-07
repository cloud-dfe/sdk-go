package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type vpo struct {
	Base base
}

func Vpo(b base) vpo {

	result := vpo{Base: b}

	return result
}

func (c vpo) Cria(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/vpo/ailog", payload)

	return resp, err
}

func (c vpo) Consulta(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/vpo/ailog/%s", key), nil)

	return resp, err
}

func (c vpo) Busca(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/vpo/ailog/busca", payload)

	return resp, err
}

func (c vpo) Cancela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/vpo/ailog/cancela", payload)

	return resp, err
}

func (c vpo) Reemite(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/vpo/ailog/reemite", payload)

	return resp, err
}

func (c vpo) Rota(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/vpo/ailog/rota", payload)

	return resp, err
}

func (c vpo) Pdf(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/vpo/ailog/pdf/%s", key), nil)

	return resp, err
}

func (c vpo) Veiculo(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/vpo/ailog/veiculo", payload)

	return resp, err
}
