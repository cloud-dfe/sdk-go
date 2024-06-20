# SDK em Golang para API Integra Notas

Este SDK visa simplificar a integração do seu sistema com a nossa API, funções para acessar as rotas da API. Isso elimina a necessidade de desenvolver uma aplicação para se comunicar diretamente com a nossa API, tornando o processo mais eficiente e direto.

*Nota: Utilizamos apenas recursos nativos do Golang para desernvolver o sdk*

## Forma de instalação de nosso SDK:

```
go get github.com/cloud-dfe/sdk-go
```

## Forma de uso:

```go
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

	nfe := sdk_cloud_dfe.Nfe(config)

	resp, err := nfe.Status()

	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Erro ao converter mapa para JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}

```

### Sobre dados de envio e retornos

Para saber os detalhes referente ao dados de envio e os retornos consulte nossa documentação [IntegraNotas Documentação](https://integranotas.com.br/doc).

### Veja alguns exemplos de consumi de nossa API nos link abaixo:

[Pasta de Exemplos](https://github.com/cloud-dfe/sdk-go/tree/master/examples)

[Utilitários](https://github.com/cloud-dfe/sdk-go/tree/master/examples/utils)

[Averbação](https://github.com/cloud-dfe/sdk-go/tree/master/examples/averbacao)

[Certificado Digital](https://github.com/cloud-dfe/sdk-go/tree/master/examples/certificado)

[CT-e](https://github.com/cloud-dfe/sdk-go/tree/master/examples/cte)

[CT-e OS](https://github.com/cloud-dfe/sdk-go/tree/master/examples/cteos)

[DF-e](https://github.com/cloud-dfe/sdk-go/tree/master/examples/dfe)

[Emitente](https://github.com/cloud-dfe/sdk-go/tree/master/examples/emitente)

[GNR-e](https://github.com/cloud-dfe/sdk-go/tree/master/examples/gnre)

[MDF-e](https://github.com/cloud-dfe/sdk-go/tree/master/examples/mdfe)

[NFC-e](https://github.com/cloud-dfe/sdk-go/tree/master/examples/nfce)

[NF-e](https://github.com/cloud-dfe/sdk-go/tree/master/examples/nfe)

[NFS-e](https://github.com/cloud-dfe/sdk-go/tree/master/examples/nfse)

[Softhouse](https://github.com/cloud-dfe/sdk-go/tree/master/examples/softhouse)