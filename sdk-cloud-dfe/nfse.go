package sdk_cloud_dfe

import (
	"errors"
	"fmt"
	"net/http"
)

type nfse struct {
	Base base
}

func Nfse(b base) nfse {

	result := nfse{Base: b}

	return result
}

func (c nfse) Cria(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfse", payload)

	return resp, err
}

func (c nfse) Preview(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfse/preview", payload)

	return resp, err
}

func (c nfse) Pdf(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfse/pdf/%s", key), nil)

	return resp, err
}

func (c nfse) Consulta(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfse/%s", key), nil)

	return resp, err
}

func (c nfse) Cancela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfse/cancela", payload)

	return resp, err
}

func (c nfse) Substitui(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfse/substitui", payload)

	return resp, err
}

func (c nfse) Busca(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfse/busca", payload)

	return resp, err
}

func (c nfse) Backup(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfse/backup", payload)

	return resp, err
}

func (c nfse) Localiza(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfse/consulta", payload)

	return resp, err
}

func (c nfse) Info(payload map[string]interface{}) (map[string]interface{}, error) {
	key, ok := payload["ibge"].(string)

	if !ok {
		return nil, errors.New("não foi informado ibge")
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfse/info/%s", key), nil)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c nfse) Conflito(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfse/conflito", payload)

	return resp, err
}

func (c nfse) Offline() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/nfse/offline", nil)

	return resp, err
}

func (c nfse) Resolve(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfse/resolve/%s", key), nil)

	return resp, err
}
