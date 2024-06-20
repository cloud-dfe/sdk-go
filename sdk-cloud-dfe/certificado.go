package sdk_cloud_dfe

type certificado struct {
	Base base
}

func Certificado(b base) certificado {

	result := certificado{Base: b}

	return result
}
