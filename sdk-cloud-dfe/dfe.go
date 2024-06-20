package sdk_cloud_dfe

type dfe struct {
	Base base
}

func Dfe(b base) dfe {

	result := dfe{Base: b}

	return result
}
