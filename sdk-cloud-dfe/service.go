package sdk_cloud_dfe

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type service struct {
	Config configService
}

func (s service) request(method, route string, payload map[string]interface{}) (interface{}, error) {

	headers := map[string]string{
		"Authorization": s.Config.Token,
		"Accept":        "application/json",
		"Content-Type":  "application/json",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		if s.Config.Debug {
			log.Fatalf("Erro ao converter dados para JSON: %v", err)
		}
		return nil, errors.New("erro ao converter dados para JSON")
	}

	url := string(s.Config.BaseUri) + route

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		if s.Config.Debug {
			log.Fatalf("Erro ao criar a requisição: %v", err)
		}
		return nil, errors.New("erro ao criar a requisição")
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(s.Config.Timeout),
	}

	resp, err := client.Do(req)
	if err != nil {
		if s.Config.Debug {
			log.Fatalf("Erro ao enviar a requisição: %v", err)
		}
		return nil, errors.New("erro ao enviar a requisição")
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		if s.Config.Debug {
			log.Fatalf("Erro ao obter a resposta: %v", err)
		}
		return nil, errors.New("erro ao obter a resposta")
	}

	return result, nil
}
