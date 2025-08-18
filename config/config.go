package config

import (
	"github.com/lengwh/dapp-wallet-sign/flags"
	"github.com/urfave/cli/v2"
)

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	LevelDBPath     string
	RpcServer       ServerConfig
	CredentialsFile string
	KeyName         string
	HsmEnable       bool
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		LevelDBPath:     ctx.String(flags.LevelDbPathFlag.Name),
		CredentialsFile: ctx.String(flags.CredentialsFileFlag.Name),
		KeyName:         ctx.String(flags.KeyNameFlag.Name),
		HsmEnable:       ctx.Bool(flags.HsmEnable.Name),
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
	}
}
