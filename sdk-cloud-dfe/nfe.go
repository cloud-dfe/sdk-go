package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type nfe struct {
	Base base
}

func Nfe(b base) nfe {

	result := nfe{Base: b}

	return result
}

func (c nfe) Cria(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe", payload)

	return resp, err
}

func (c nfe) Preview(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/preview", payload)

	return resp, err
}

func (n nfe) Status() (interface{}, error) {
	resp, err := n.Base.Client.send(http.MethodGet, "/nfe/status", nil)

	return resp, err
}

func (c nfe) Consulta(payload map[string]interface{}) (interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfe/%s", key), nil)

	return resp, err
}

func (c nfe) Busca(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/busca", payload)

	return resp, err
}

func (c nfe) Cancela(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/cancela", payload)

	return resp, err
}

func (c nfe) Correcao(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/correcao", payload)

	return resp, err
}

func (c nfe) Inutiliza(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/inutiliza", payload)

	return resp, err
}

func (c nfe) Pdf(payload map[string]interface{}) (interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfe/pdf/%s", key), nil)

	return resp, err
}

func (c nfe) Etiqueta(payload map[string]interface{}) (interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/nfe/etiqueta/%s", key), nil)

	return resp, err
}

func (c nfe) Manifesta(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/manifesta", payload)

	return resp, err
}

func (c nfe) Backup(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/backup", payload)

	return resp, err
}

func (c nfe) Download(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/download", payload)

	return resp, err
}

func (c nfe) Recebidas(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/recebidas", payload)

	return resp, err
}

func (c nfe) Interessado(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/interresado", payload)

	return resp, err
}

func (c nfe) Importa(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/importa", payload)

	return resp, err
}

func (c nfe) Comprovante(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/comprovante", payload)

	return resp, err
}

func (c nfe) Cadastro(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfe/cadastro", payload)

	return resp, err
}
