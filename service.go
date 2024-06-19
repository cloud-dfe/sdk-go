package sdk_cloud_dfe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type service struct {
	Config configService
}

func setService(config configService) service {

	return service{
		Config: config,
	}
}

func (s service) Request(method, route string, payload interface{}) (interface{}, error) {

	headers := map[string]string{
		"Authorization": s.Config.Token,
		"Accept":        "application/json",
		"Content-Type":  "application/json",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		if s.Config.Debug {
			fmt.Printf("Erro ao converter dados para JSON: %v", err)
		}
		return nil, err
	}

	url := string(s.Config.BaseUri) + route

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		if s.Config.Debug {
			fmt.Printf("Erro ao criar a requisição: %v", err)
		}
		return nil, err
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
			fmt.Printf("Erro ao enviar a requisição: %v", err)
		}
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		if s.Config.Debug {
			fmt.Printf("Erro ao obter a resposta: %v", err)
		}
		return nil, err
	}

	return result, nil
}
