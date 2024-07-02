package main

import (
	"encoding/json"
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOiJ0b2tlbl9leGVtcGxvIiwidXNyIjoidGsiLCJ0cCI6InRrIn0.Tva_viCMCeG3nkRYmi_RcJ6BtSzui60kdzIsuq5X-sQ"

	config, err := sdk_cloud_dfe.NewBase(token, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	nfe := sdk_cloud_dfe.Nfe(config)

	payload := map[string]interface{}{
		"chave": "50000000000000000000000000000000000000000000", // Obrigatoria Chave de acesso
		"registra": map[string]interface{}{ // dados opcionais no caso de cancelamento
			"data":                "2021-10-12T12:22:33-03:00", // Obrigatório Data e Hora do recebimento. (dhEntrega)
			
			
			"imagem":              "lUHJvYyB2ZXJzYW....",       // Opcional Base64 da imagem capturada do Comprovante de Entrega da nNF-e ou uma string de referencia
			"recebedor_documento": "123456789 SSPRJ",           // Obrigatório Número do documento de identificação da pessoa que assinou o Comprovante de Entrega da NF-e. (nDoc)
			"recebedor_nome":      "NOME TESTE",                // Obrigatório Nome da pessoa que assinou o Comprovante de Entrega da NF-e. (xNome)
			"coordenadas": map[string]interface{}{ // dados opcionais no caso de cancelamento
				"latitude":  -23.628360, // Latitude do ponto de entrega, com 6 decimais. (latGPS)
				"longitude": -46.622109, // Longitude do ponto de entrega, com 6 decimais. (longGPS)
			},
		},
	}

	resp, err := nfe.Comprovante(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
