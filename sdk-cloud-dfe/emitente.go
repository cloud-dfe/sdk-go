package sdk_cloud_dfe

import "net/http"

type emitente struct {
	Base base
}

func Emitente(b base) emitente {

	result := emitente{Base: b}

	return result
}

func (c emitente) Token() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/emitente/token", nil)

	return resp, err
}

func (c emitente) Atualiza(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPut, "/emitente", payload)

	return resp, err
}

func (c emitente) Mostra() (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/emitente", nil)

	return resp, err
}
