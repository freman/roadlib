package speed

type (
	KMHR float64
	MHR  float64
	MM   float64
	MS   float64
)

func (k KMHR) MetresHr() MHR {
	return MHR(k * 1000)
}

func (k KMHR) MetresMinute() MM {
	return MM(k * 1000 / 60)
}

func (k KMHR) MetresSecond() MS {
	return MS(k * 1000 / 60 / 60)
}

func (m MHR) MetresMinute() MM {
	return MM(m / 60)
}

func (m MM) MetresSecond() MS {
	return MS(m / 60)
}
