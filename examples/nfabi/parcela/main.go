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

	nfabi := sdk_cloud_dfe.Nfabi(config)

	var payload map[string]interface{}
	json.Unmarshal([]byte(`{
	    "chave": "41261210944071000104770010000001231012345679",
	    "codigo": "01",
	    "valor_parcela": "10000.00",
	    "valor_total_parcela": "10000.00",
	    "data_pagamento": "2027-01-10",
	    "tributacao": {
	        "cst": "200",
	        "classificacao_tributaria": "200046",
	        "valor_operacao": "10000.00",
	        "base_calculo": "10000.00",
	        "ibs_uf": {
	            "aliquota": "0.1000",
	            "reducao_aliquota": "50.0000",
	            "aliquota_efetiva": "0.0500",
	            "valor": "5.00"
	        },
	        "ibs_municipal": {
	            "aliquota": "0.0000",
	            "reducao_aliquota": "50.0000",
	            "aliquota_efetiva": "0.0000",
	            "valor": "0.00"
	        },
	        "cbs": {
	            "aliquota": "0.9000",
	            "reducao_aliquota": "50.0000",
	            "aliquota_efetiva": "0.4500",
	            "valor": "45.00"
	        }
	    }
	}`), &payload)

	resp, err := nfabi.Parcela(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
