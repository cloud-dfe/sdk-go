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

	nfce := sdk_cloud_dfe.Nfce(config)

	listaItens := []map[string]interface{}{
		{
			"numero_item":               "1",
			"codigo_produto":            "000297",
			"descricao":                 "SAL GROSSO 50KGS",
			"codigo_ncm":                "84159020",
			"cfop":                      "5102",
			"unidade_comercial":         "SC",
			"quantidade_comercial":      10,
			"valor_unitario_comercial":  "22.45",
			"valor_bruto":               "224.50",
			"unidade_tributavel":        "SC",
			"quantidade_tributavel":     "10.00",
			"valor_unitario_tributavel": "22.45",
			"origem":                    "0",
			"inclui_no_total":           "1",
			"imposto": map[string]interface{}{
				"valor_aproximado_tributos": 9.43,
				"icms": map[string]interface{}{
					"situacao_tributaria":                 "102",
					"aliquota_credito_simples":            "0",
					"valor_credito_simples":               "0",
					"modalidade_base_calculo":             "3",
					"valor_base_calculo":                  "0.00",
					"modalidade_base_calculo_st":          "4",
					"aliquota_reducao_base_calculo":       "0.00",
					"aliquota":                            "0.00",
					"aliquota_final":                      "0.00",
					"valor":                               "0.00",
					"aliquota_margem_valor_adicionado_st": "0.00",
					"aliquota_reducao_base_calculo_st":    "0.00",
					"valor_base_calculo_st":               "0.00",
					"aliquota_st":                         "0.00",
					"valor_st":                            "0.00",
				},
				"fcp": map[string]interface{}{
					"aliquota": "1.65",
				},
				"pis": map[string]interface{}{
					"situacao_tributaria": "01",
					"valor_base_calculo":  224.5,
					"aliquota":            "1.65",
					"valor":               "3.70",
				},
				"cofins": map[string]interface{}{
					"situacao_tributaria": "01",
					"valor_base_calculo":  224.5,
					"aliquota":            "7.60",
					"valor":               "17.06",
				},
			},
			"valor_desconto":              0,
			"valor_frete":                 0,
			"valor_seguro":                0,
			"valor_outras_despesas":       0,
			"informacoes_adicionais_item": "Valor aproximado tributos R$: 9,43 (4,20%) Fonte: IBPT",
		},
	}

	payload := map[string]interface{}{
		"natureza_operacao":  "VENDA DENTRO DO ESTADO",
		"serie":              "1",
		"numero":             "101008",
		"data_emissao":       "2021-06-26T15:20:00-03:00",
		"tipo_operacao":      "1",
		"presenca_comprador": "1",
		"itens":              listaItens,
		"frete": map[string]interface{}{
			"modalidade_frete": "9",
		},
		"pagamento": map[string]interface{}{
			"formas_pagamento": []map[string]interface{}{
				{
					"meio_pagamento": "01",
					"valor":          "224.50",
				},
			},
		},
		"informacoes_adicionais_contribuinte": "PV: 3325 * Rep: DIRETO * Motorista:  * Forma Pagto: 04 DIAS * teste de observação para a nota fiscal * Valor aproximado tributos R$9,43 (4,20%) Fonte: IBPT",
		"pessoas_autorizadas": []map[string]interface{}{
			{
				"cnpj": "96256273000170",
			},
			{
				"cnpj": "80681257000195",
			},
		},
	}

	resp, err := nfce.Preview(payload)

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
