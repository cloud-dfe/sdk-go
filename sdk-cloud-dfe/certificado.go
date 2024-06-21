package sdk_cloud_dfe

import "net/http"

type certificado struct {
	Base base
}

func Certificado(b base) certificado {

	result := certificado{Base: b}

	return result
}

func (c certificado) Atualiza(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/certificado", payload)

	return resp, err
}

func (c certificado) Mostra() (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodGet, "/certificado", nil)

	return resp, err
}
