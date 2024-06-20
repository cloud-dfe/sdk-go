package sdk_cloud_dfe

type cte struct {
	Base base
}

func Cte(b base) cte {

	result := cte{Base: b}

	return result
}
