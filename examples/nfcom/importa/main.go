package main

import (
	"encoding/json"
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOjQ2MSwidXNyIjoxNzAsInRwIjoyLCJpYXQiOjE2NTE1MDYzMjR9.a0cOwP6BUDZAboYwMzoMjutCtFM8Ph-X4pLahZIB_V4"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	nfcom := sdk_cloud_dfe.Nfcom(config)

	xml, err := sdk_cloud_dfe.ReadFile("caminho_do_arquivo.xml")

	if err != nil {
		fmt.Println("Ocorreu um erro ao tentar ler o arquivo XML")
	}

	fileXmlBase64 := sdk_cloud_dfe.Encode(xml)

	payload := map[string]interface{}{
		"xml": fileXmlBase64,
	}

	resp, err := nfcom.Importa(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
