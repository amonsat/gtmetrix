package gtmetrix

type ThrottleType string

const (
	ThrottleTypeBroadbandFast = "20000/5000/25"
	ThrottleTypeBroadband     = "5000/1000/30"
	ThrottleTypeBroadbandSlow = "1500/384/50"
	ThrottleTypeLteMobile     = "15000/10000/100"
	ThrottleType3gMobile      = "1600/768/200"
	ThrottleType2gMobile      = "240/200/400"
	ThrottleType56kDialUp     = "50/30/125"
)
