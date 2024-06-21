package sdk_cloud_dfe

type client struct {
	Service service
}

func (c client) send(method, route string, payload map[string]interface{}) (interface{}, error) {

	send, err := c.Service.request(method, route, payload)

	return send, err

}
