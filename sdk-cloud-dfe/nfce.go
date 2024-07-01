package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type nfce struct {
	Base base
}

func Nfce(b base) nfce {

	result := nfce{Base: b}

	return result
}

func (c nfce) Cria(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfce", payload)

	return resp, err
}

func (c nfce) Preview(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfce/preview", payload)

	return resp, err
}

func (c nfce) Status() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/nfce/status", nil)

	return resp, err
}

func (c nfce) Consulta(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfce/%s", key), nil)

	return resp, err
}

func (c nfce) Busca(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfce/busca", payload)

	return resp, err
}

func (c nfce) Cancela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfce/cancela", payload)

	return resp, err
}

func (c nfce) Offline() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/nfce/offline", nil)

	return resp, err
}

func (c nfce) Inutiliza(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfce/inutiliza", payload)

	return resp, err
}

func (c nfce) Pdf(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfce/pdf/%s", key), nil)

	return resp, err
}

func (c nfce) Substitui(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfce/substitui", payload)

	return resp, err
}

func (c nfce) Backup(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfce/backup", payload)

	return resp, err
}

func (c nfce) Importa(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfce/importa", payload)

	return resp, err
}
