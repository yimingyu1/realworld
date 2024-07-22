package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Redis struct {
		Host string
		Type string
	}

	DB struct {
		DataSource string
	}

	JwtAuth struct {
		AccessSecret string
		AccessExpire int
	}
}
