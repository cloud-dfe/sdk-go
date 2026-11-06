package sdk_cloud_dfe

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type ciot struct {
	Base base
}

func Ciot(b base) ciot {

	result := ciot{Base: b}

	return result
}

// o CIOT é identificado pela identificação informada na criação, não por chave de acesso
var identificacaoCiot = regexp.MustCompile(`^[A-Za-z0-9._-]{1,40}$`)

func checkIdentificacao(payload map[string]interface{}) (string, error) {
	identificacao, _ := payload["identificacao"].(string)
	identificacao = strings.TrimSpace(identificacao)

	if !identificacaoCiot.MatchString(identificacao) {
		return "", errors.New("a identificação do CIOT deve ter de 1 a 40 caracteres: letras, números, ponto, hífen ou sublinhado")
	}

	return url.PathEscape(identificacao), nil
}

func (c ciot) Cria(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot", payload)

	return resp, err
}

func (c ciot) Consulta(payload map[string]interface{}) (map[string]interface{}, error) {
	identificacao, err := checkIdentificacao(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/ciot/%s", identificacao), nil)

	return resp, err
}

func (c ciot) Busca(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/busca", payload)

	return resp, err
}

func (c ciot) Cancela(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/cancela", payload)

	return resp, err
}

func (c ciot) Encerra(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/encerra", payload)

	return resp, err
}

func (c ciot) Retifica(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/retifica", payload)

	return resp, err
}

func (c ciot) Pdf(payload map[string]interface{}) (map[string]interface{}, error) {
	identificacao, err := checkIdentificacao(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.Client.send(http.MethodGet, fmt.Sprintf("/ciot/pdf/%s", identificacao), nil)

	return resp, err
}

func (c ciot) Transportador(payload map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.Base.Client.send(http.MethodPost, "/ciot/transportador", payload)

	return resp, err
}
