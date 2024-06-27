package sdk_cloud_dfe

import "net/http"

type averbacao struct {
	Base base
}

func Averbacao(b base) averbacao {

	result := averbacao{Base: b}

	return result
}

func (a averbacao) Atm(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := a.Base.Client.send(http.MethodPost, "/averbacao/atm", payload)

	return resp, err
}

func (a averbacao) AtmCancela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := a.Base.Client.send(http.MethodPost, "/averbacao/atm/cancela", payload)

	return resp, err
}

func (a averbacao) Elt(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := a.Base.Client.send(http.MethodPost, "/averbacao/elt", payload)

	return resp, err
}

func (a averbacao) PortoSeguro(payload map[string]interface{}) (interface{}, error) {
	resp, err := a.Base.Client.send(http.MethodPost, "/averbacao/portoseguro", payload)

	return resp, err
}

func (a averbacao) PortoSeguroCancela(payload map[string]interface{}) (interface{}, error) {
	resp, err := a.Base.Client.send(http.MethodPost, "/averbacao/portoseguro/cancela", payload)

	return resp, err
}
