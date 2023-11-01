package config

type Frp struct {
	bindAddr    string
	bindPort    int64
	kcpBindPort int64
}

var FrpConfig = new(Frp)
