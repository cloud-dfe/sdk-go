package sdk_cloud_dfe

import "fmt"

type ambiente int
type baseUri string

const (
	AmbienteProducao    ambiente = 1
	AmbienteHomologacao ambiente = 2
	api_Producao        baseUri  = "https://api.integranotas.com.br/v1"
	api_Homologacao     baseUri  = "https://hom-api.integranotas.com.br/v1"
)

type configBase struct {
	Token    string
	Ambiente ambiente
	Timeout  int
	Port     int
	Debug    bool
}

func NewBase(token string, ambiente ambiente, timeout int, port int, debug bool) configBase {

	config := configBase{
		Token:    token,
		Ambiente: ambiente,
		Timeout:  timeout,
		Port:     port,
		Debug:    debug,
	}

	if debug {
		printBase(config)
	}

	return config
}

type configClient struct {
	Token    string
	Ambiente ambiente
	Timeout  int
	Port     int
	Debug    bool
}

func NewClient(config configBase) configClient {

	token := config.Token
	ambiente := config.Ambiente
	timeout := config.Timeout
	port := config.Port
	debug := config.Debug

	client := configClient{
		Token:    token,
		Ambiente: ambiente,
		Timeout:  timeout,
		Port:     port,
		Debug:    debug,
	}

	if debug {
		printClient(client)
	}

	return client
}

type configService struct {
	BaseUri baseUri
	Token   string
	Timeout int
	Port    int
	Debug   bool
}

func NewService(config configClient) (service, error) {
	token := config.Token
	ambiente := config.Ambiente
	timeout := config.Timeout
	port := config.Port
	debug := config.Debug

	if ambiente != AmbienteProducao && ambiente != AmbienteHomologacao {
		return service{}, nil

	} else {
		if ambiente == AmbienteProducao {
			baseUri := api_Producao

			newService := configService{
				BaseUri: baseUri,
				Token:   token,
				Timeout: timeout,
				Port:    port,
				Debug:   debug,
			}

			if debug {
				printRequest(newService)
			}

			service := setService(newService)

			return service, nil

		} else {
			baseUri := api_Homologacao

			newService := configService{
				BaseUri: baseUri,
				Token:   token,
				Timeout: timeout,
				Port:    port,
				Debug:   debug,
			}

			if debug {
				printRequest(newService)
			}

			service := setService(newService)

			return service, nil
		}
	}
}

func printBase(config configBase) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Ambiente: %d\n", config.Ambiente)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
	fmt.Println("")
}

func printClient(config configClient) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Ambiente: %d\n", config.Ambiente)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
	fmt.Println("")
}

func printRequest(config configService) {
	fmt.Printf("Token: %s\n", config.Token)
	fmt.Printf("Base URI: %s\n", config.BaseUri)
	fmt.Printf("Timeout: %d\n", config.Timeout)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
	fmt.Println("")
}
