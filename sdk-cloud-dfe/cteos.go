package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type cteos struct {
	Base base
}

func Cteos(b base) cteos {

	result := cteos{Base: b}

	return result
}

func (c cteos) Status(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/status", payload)

	return resp, err
}

func (c cteos) Consulta(payload map[string]interface{}) (interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/cteos/%s", key), nil)

	return resp, err
}

func (c cteos) Pdf(payload map[string]interface{}) (interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/cteos/pdf/%s", key), nil)

	return resp, err
}

func (c cteos) Cria(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos", payload)

	return resp, err
}

func (c cteos) Busca(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/busca", payload)

	return resp, err
}

func (c cteos) Cancela(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/cancela", payload)

	return resp, err
}

func (c cteos) Correcao(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/correcao", payload)

	return resp, err
}

func (c cteos) Inutiliza(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/inutiliza", payload)

	return resp, err
}

func (c cteos) Backup(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/backup", payload)

	return resp, err
}

func (c cteos) Importa(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/importa", payload)

	return resp, err
}

func (c cteos) Preview(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/preview", payload)

	return resp, err
}

func (c cteos) Desacordo(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/cteos/desacordo", payload)

	return resp, err
}
