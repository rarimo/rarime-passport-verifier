package vault

import (
	"context"
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/rarimo/rarime-passport-verifier/internal/config"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
)

const (
	vaultIssuerPath   = "issuer"
	vaultVerifierPath = "verifier"
)

type VaultClient struct {
	client    *vaultapi.Client
	mountPath string
}

func NewVaultClient(cfg *config.VaultConfig) (*VaultClient, error) {
	conf := vaultapi.DefaultConfig()
	conf.Address = cfg.Address

	client, err := vaultapi.NewClient(conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize new client")
	}

	client.SetToken(cfg.Token)

	return &VaultClient{
		client:    client,
		mountPath: cfg.MountPath,
	}, nil
}

func (v *VaultClient) IssuerAuthData() (string, string, error) {
	conf := struct {
		IssuerLogin    string `fig:"login,required"`
		IssuerPassword string `fig:"password,required"`
	}{}

	secret, err := v.client.KVv2(v.mountPath).Get(context.Background(), vaultIssuerPath)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to get secret")
	}

	if err := figure.
		Out(&conf).
		With(figure.BaseHooks).
		From(secret.Data).
		Please(); err != nil {
		return "", "", errors.Wrap(err, "failed to figure out")
	}

	return conf.IssuerLogin, conf.IssuerPassword, nil
}

func (v *VaultClient) Blinder() (*big.Int, error) {
	conf := struct {
		Blinder string `fig:"blinder,required"`
	}{}

	secret, err := v.client.KVv2(v.mountPath).Get(context.Background(), vaultVerifierPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get secret")
	}

	if err := figure.
		Out(&conf).
		With(figure.BaseHooks).
		From(secret.Data).
		Please(); err != nil {
		return nil, errors.Wrap(err, "failed to figure out")
	}

	blinder, ok := new(big.Int).SetString(conf.Blinder, 10)
	if !ok {
		return nil, errors.New("failed to set string to big.Int")
	}

	return blinder, nil
}
