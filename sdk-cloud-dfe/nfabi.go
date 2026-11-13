package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type nfabi struct {
	Base base
}

func Nfabi(b base) nfabi {

	result := nfabi{Base: b}

	return result
}

func (c nfabi) Status() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/nfabi/status", nil)

	return resp, err
}

func (c nfabi) Cria(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfabi", payload)

	return resp, err
}

func (c nfabi) Consulta(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfabi/%s", key), nil)

	return resp, err
}

func (c nfabi) Cancela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfabi/cancela", payload)

	return resp, err
}

func (c nfabi) Parcela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfabi/parcela", payload)

	return resp, err
}

func (c nfabi) Desfazimento(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfabi/desfazimento", payload)

	return resp, err
}

func (c nfabi) Busca(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfabi/busca", payload)

	return resp, err
}

func (c nfabi) Pdf(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfabi/pdf/%s", key), nil)

	return resp, err
}

func (c nfabi) Preview(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfabi/preview", payload)

	return resp, err
}

func (c nfabi) Backup(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfabi/backup", payload)

	return resp, err
}

func (c nfabi) Importa(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfabi/importa", payload)

	return resp, err
}
