package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type nfcom struct {
	Base base
}

func Nfcom(b base) nfcom {

	result := nfcom{Base: b}

	return result
}

func (c nfcom) Status() (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/nfcom/status", nil)

	return resp, err
}

func (c nfcom) Cria(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfcom", payload)

	return resp, err
}

func (c nfcom) Consulta(payload map[string]interface{}) (interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfcom/%s", key), nil)

	return resp, err
}

func (c nfcom) Cancela(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfcom/cancela", payload)

	return resp, err
}

func (c nfcom) Busca(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfcom/busca", payload)

	return resp, err
}

func (c nfcom) Pdf(payload map[string]interface{}) (interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfcom/pdf/%s", key), nil)

	return resp, err
}

func (c nfcom) Preview(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfcom/preview", payload)

	return resp, err
}

func (c nfcom) Backup(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfcom/backup", payload)

	return resp, err
}

func (c nfcom) Importa(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfcom/importa", payload)

	return resp, err
}
