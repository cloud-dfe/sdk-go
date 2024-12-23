package main

import (
	"encoding/json"
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	terren := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbXAiOiJ0b2tlbl9leGVtcGxvIiwidXNyIjoidGsiLCJ0cCI6InRrIn0.Tva_viCMCeG3nkRYmi_RcJ6BtSzui60kdzIsuq5X-sQ"

	config, err := sdk_cloud_dfe.NewBase(terren, sdk_cloud_dfe.AmbienteHomologacao, 60, 443, false)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	mdfe := sdk_cloud_dfe.Mdfe(config)

	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo.")
	}

	payload := map[string]interface{}{
		"tipo_operacao":   "2",
		"tipo_transporte": nil,
		"numero":          "27",
		"serie":           "1",
		"data_emissao":    "2021-06-26T09:21:42-00:00",
		"uf_inicio":       "RN",
		"uf_fim":          "GO",
		"municipios_carregamento": []map[string]interface{}{
			{
				"codigo_municipio": "2408003",
				"nome_municipio":   "Mossoró",
			},
		},
		"percursos": []map[string]interface{}{
			{"uf": "PB"},
			{"uf": "PE"},
			{"uf": "BA"},
		},
		"municipios_descarregamento": []map[string]interface{}{
			{
				"codigo_municipio": "5200050",
				"nome_municipio":   "Abadia de Goiás",
				"nfes": []map[string]interface{}{
					{
						"chave": "50000000000000000000000000000000000000000000",
					},
				},
			},
		},
		"valores": map[string]interface{}{
			"valor_total_carga":                "100.00",
			"codigo_unidade_medida_peso_bruto": "01",
			"peso_bruto":                       "1000.000",
		},
		"informacao_adicional_fisco": nil,
		"informacao_complementar":    nil,
		"modal_rodoviario": map[string]interface{}{
			"rntrc":        "57838055",
			"ciot":         nil,
			"contratante":  nil,
			"vale_pedagio": nil,
			"veiculo": map[string]interface{}{
				"codigo":          "000000001",
				"placa":           "FFF1257",
				"renavam":         "335540391",
				"tara":            "0",
				"tipo_rodado":     "01",
				"tipo_carroceria": "00",
				"uf":              "MT",
				"condutores": []map[string]interface{}{
					{
						"nome": "JOAO TESTE",
						"cpf":  "01234567890",
					},
				},
			},
			"reboques": nil,
		},
	}

	resp, err := mdfe.Cria(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

	if sucesso, err := resp["sucesso"].(bool); err && sucesso {
		if codigo, err := resp["codigo"].(float64); err && codigo == 2 {
			// Offline
			fmt.Println("Documento offline. Aguarde a notificação.")
		} else {
			// Autorizado
			fmt.Printf("Documento autorizado: %+v\n", resp)
		}
	} else if codigo, err := resp["codigo"].(float64); err && (codigo == 5001 || codigo == 5002) {
		// Erro nos campos
		if erros, err := resp["erros"]; err {
			fmt.Printf("Erro nos campos: %+v\n", erros)
		} else {
			fmt.Println("Erro nos campos, mas sem detalhes disponíveis.")
		}
	} else if codigo, err := resp["codigo"].(float64); err && (codigo == 5008 || codigo >= 7000) {
		chave, chaveerr := resp["chave"].(string)
		if !chaveerr {
			fmt.Println("Chave não encontrada no response.")
			return
		}
	
		// >= 7000 indica problemas de comunicação com a SEFAZ
		fmt.Printf("Problemas de comunicação ou chave pendente: %+v\n", resp)
	
		payloadConsulta := map[string]interface{}{
			"chave": chave,
		}
	
		respConsulta, err := mdfe.Consulta(payloadConsulta)
		if err != nil {
			fmt.Printf("Erro ao consultar documento: %v\n", err)
			return
		}
	
		if consultaCodigo, err := respConsulta["codigo"].(float64); err && consultaCodigo != 5023 {
			if consultaSucesso, err := respConsulta["sucesso"].(bool); err && consultaSucesso {
				// Autorizado
				fmt.Printf("Documento autorizado após consulta: %+v\n", respConsulta)
			} else {
				// Rejeição
				fmt.Printf("Documento rejeitado após consulta: %+v\n", respConsulta)
			}
		} else {
			// Em processamento
			fmt.Printf("Documento em processamento: %+v\n", respConsulta)
		}
	} else {
		// Rejeição
		fmt.Printf("Documento rejeitado: %+v\n", resp)
	}

}
