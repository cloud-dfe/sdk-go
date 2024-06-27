package sdk_cloud_dfe

import (
	"fmt"
	"net/http"
)

type dfe struct {
	Base base
}

func Dfe(b base) dfe {

	result := dfe{Base: b}

	return result
}

func (c dfe) BuscaCte(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/dfe/cte", payload)

	return resp, err
}

func (c dfe) DownloadCte(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/dfe/cte/%s", key), nil)

	return resp, err
}

func (c dfe) BuscaNfe(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/dfe/nfe", payload)

	return resp, err
}

func (c dfe) DownloadNfe(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/dfe/nfe/%s", key), nil)

	return resp, err
}

func (c dfe) BuscaNfse(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/dfe/nfse", payload)

	return resp, err
}

func (c dfe) DownloadNfse(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/dfe/nfse/%s", key), nil)

	return resp, err
}

func (c dfe) Eventos(payload map[string]interface{}) (map[string]interface{}, error) {
	key, err := checkKey(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/dfe/eventos/%s", key), nil)

	return resp, err
}

func (c dfe) Backup(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/dfe/backup", payload)

	return resp, err
}
