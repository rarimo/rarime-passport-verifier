package config

import (
	points "github.com/rarimo/rarime-points-svc/pkg/connector"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	IssuerConfiger
	VerifierConfiger
	NetworkConfiger
	VaultConfiger
	points.Pointer
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	getter kv.Getter

	IssuerConfiger
	VerifierConfiger
	NetworkConfiger
	VaultConfiger
	points.Pointer
}

func New(getter kv.Getter) Config {
	return &config{
		getter:           getter,
		Databaser:        pgdb.NewDatabaser(getter),
		Copuser:          copus.NewCopuser(getter),
		Listenerer:       comfig.NewListenerer(getter),
		Logger:           comfig.NewLogger(getter, comfig.LoggerOpts{}),
		IssuerConfiger:   NewIssuerConfiger(getter),
		VerifierConfiger: NewVerifierConfiger(getter),
		NetworkConfiger:  NewNetworkConfiger(getter),
		VaultConfiger:    NewVaultConfiger(getter),
		Pointer:          points.NewPointer(getter),
	}
}
