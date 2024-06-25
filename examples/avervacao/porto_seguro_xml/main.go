package main

import (
	"encoding/json"
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOjE0LCJ1c3IiOjgsInRwIjoyLCJpYXQiOjE2NzIyNTAzMzV9.TY8-eAg6gUFSo55epFL-UoPTD3XAUJMl8SxUcAsCr4o"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	averbacao := sdk_cloud_dfe.Averbacao(config)

	fileXmlBase64, err := sdk_cloud_dfe.ReadFile("sdk-go/caminho_do_arquivo.xml")

	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo.")
	}

	payload := map[string]interface{}{
		"xml":     fileXmlBase64,
		"usuario": "login",
		"senha":   "senha",
		"chave":   "50000000000000000000000000000000000000000000",
	}

	resp, err := averbacao.PortoSeguro(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
