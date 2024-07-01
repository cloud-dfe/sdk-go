package main

import (
	"fmt"

	sdk_cloud_dfe "github.com/cloud-dfe/sdk-go/sdk-cloud-dfe"
)

func main() {

	encodedContent := sdk_cloud_dfe.Encode("Texto para decodificar")
	fmt.Println("\nEncoded :", encodedContent)

}
