package main

import (
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go"
)

func main() {

	config := sdk_cloud_dfe.NewBase("Teste", sdk_cloud_dfe.AMBIENTE_HOMOLOGACAO, 60, 443, false)
	config2 := sdk_cloud_dfe.NewClient(config)
	config3, err := sdk_cloud_dfe.NewRequest(config2)

	if err != nil {
		println("Ocorreu um erro")
	}

	sdk_cloud_dfe.PrintConfig(config)
	fmt.Println("")
	sdk_cloud_dfe.PrintConfig2(config2)
	fmt.Println("")
	sdk_cloud_dfe.PrintConfig3(config3)
}
