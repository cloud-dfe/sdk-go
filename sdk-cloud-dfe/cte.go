package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type cte struct {
	Base base
}

func Cte(b base) cte {

	result := cte{Base: b}

	return result
}

func (c cte) Status() (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/cte/status", nil)

	return resp, err
}

func (c cte) Consulta(payload map[string]interface{}) (interface{}, error) {

	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/cte/%s", key), nil)

	return resp, err
}

func (c cte) Pdf(payload map[string]interface{}) (interface{}, error) {

	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/cte/pdf/%s", key), nil)

	return resp, err
}

func (c cte) Cria(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte", payload)

	return resp, err
}

func (c cte) Busca(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte/busca", payload)

	return resp, err
}

func (c cte) Cancela(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte/cancela", payload)

	return resp, err
}

func (c cte) Correcao(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte/correcao", payload)

	return resp, err
}

func (c cte) Inutiliza(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte/inutiliza", payload)

	return resp, err
}

func (c cte) Backup(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte/backup", payload)

	return resp, err
}

func (c cte) Importa(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte/importa", payload)

	return resp, err
}

func (c cte) Preview(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte/preview", payload)

	return resp, err
}

func (c cte) Desacordo(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cte/desacordo", payload)

	return resp, err
}
