package sdk_cloud_dfe

import "net/http"

type nfe struct {
	Base base
}

func Nfe(b base) nfe {

	result := nfe{Base: b}

	return result
}

func (n nfe) Status() (interface{}, error) {
	resp, err := n.Base.Client.send(http.MethodGet, "/nfe/status", nil)

	return resp, err
}
