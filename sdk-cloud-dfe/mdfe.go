package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type mdfe struct {
	Base base
}

func Mdfe(b base) mdfe {

	result := mdfe{Base: b}

	return result
}

func (c mdfe) Cria(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe", payload)

	return resp, err
}

func (c mdfe) Preview(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe/preview", payload)

	return resp, err
}

func (c mdfe) Status() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/mdfe/status", nil)

	return resp, err
}

func (c mdfe) Consulta(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/mdfe/%s", key), nil)

	return resp, err
}

func (c mdfe) Busca(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe/busca", payload)

	return resp, err
}

func (c mdfe) Cancela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe/cancela", payload)

	return resp, err
}

func (c mdfe) Encerra(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe/encerra", payload)

	return resp, err
}

func (c mdfe) Condutor(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe/condutor", payload)

	return resp, err
}

func (c mdfe) Offline() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/mdfe/offline", nil)

	return resp, err
}

func (c mdfe) Pdf(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/mdfe/%s", key), nil)

	return resp, err
}

func (c mdfe) Backup(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe/backup", payload)

	return resp, err
}

func (c mdfe) Nfe(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe/nfe", payload)

	return resp, err
}

func (c mdfe) Abertos() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/mdfe/abertos", nil)

	return resp, err
}

func (c mdfe) Importa(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/mdfe/importa", payload)

	return resp, err
}
