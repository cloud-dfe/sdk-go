package sdk_cloud_dfe

import (
	"net/http"
)

type nfcom struct {
	Base base
}

func Nfcom(b base) nfcom {

	result := nfcom{Base: b}

	return result
}

func (c nfcom) Cria(payload map[string]interface{}) (interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/nfcom", payload)

	return resp, err
}
