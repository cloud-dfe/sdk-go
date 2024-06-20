package sdk_cloud_dfe

type emitente struct {
	Base base
}

func Emitente(b base) emitente {

	result := emitente{Base: b}

	return result
}
