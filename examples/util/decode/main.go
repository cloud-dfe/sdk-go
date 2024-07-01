package main

import (
	"fmt"
	"log"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {
	decodedContent, err := sdk_cloud_dfe.Decode("VGV4dG8gcGFyYSBkZWNvZGlmaWNhcg==")
	if err != nil {
		log.Fatalf("Erro ao decodificar o conteúdo: %v", err)
	}
	fmt.Println("\nDecoded :", decodedContent)
}
