package sdk_cloud_dfe

type client struct {
	Service service
}

func setClient(config configClient) (service, error) {

	service, err := NewService(config)

	return service, err

}
