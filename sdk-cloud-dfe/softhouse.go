package sdk_cloud_dfe

import (
	"errors"
	"fmt"
	"net/http"
)

type softhouse struct {
	Base base
}

func Softhouse(b base) softhouse {

	result := softhouse{Base: b}

	return result
}

func (c nfse) CriaEmitente(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/soft/emitente", payload)

	return resp, err
}

func (c nfse) AtualizaEmitente(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/soft/emitente", payload)

	return resp, err
}

func (c nfse) MostraEmitente(payload map[string]interface{}) (map[string]interface{}, error) {
	doc, ok := payload["doc"].(string)

	if !ok {
		return nil, errors.New("não foi informado doc")
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/soft/emitente/%s", doc), nil)

	return resp, err
}

func (c softhouse) ListaEmitente(payload map[string]interface{}) (map[string]interface{}, error) {

	status, ok := payload["status"]

	if !ok {
		return nil, errors.New("deve ser passado um status para listar os emitentes")
	}

	rota := "/soft/emitente"

	if status == "deletados" || status == "inativos" {
		rota = "/soft/emitente/deletados"
	}

	resp, err := c.Base.Client.send(http.MethodGet, rota, nil)

	return resp, err

}

func (c softhouse) DeletaEmitente(payload map[string]interface{}) (map[string]interface{}, error) {

	doc, ok := payload["doc"]

	if !ok {
		return nil, errors.New("deve ser passado um CNPJ ou um CPF para visualizar o emitente")
	}

	resp, err := c.Base.Client.send(http.MethodDelete, fmt.Sprintf("/soft/emitente/%s", doc), nil)

	return resp, err

}
