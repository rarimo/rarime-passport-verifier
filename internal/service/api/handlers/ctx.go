package handlers

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	stateabi "github.com/iden3/contracts-abi/state/go/abi"
	"github.com/rarimo/rarime-passport-verifier/internal/config"
	"github.com/rarimo/rarime-passport-verifier/internal/data"
	"github.com/rarimo/rarime-passport-verifier/internal/service/issuer"
	"github.com/rarimo/rarime-passport-verifier/internal/service/vault"
	points "github.com/rarimo/rarime-points-svc/pkg/connector"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	masterQKey
	verifierConfigKey
	stateContractKey
	issuerCtxKey
	vaultClientCtxKey
	ethClientCtxKey
	pointsCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxMasterQ(entry data.MasterQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, masterQKey, entry)
	}
}

func MasterQ(r *http.Request) data.MasterQ {
	return r.Context().Value(masterQKey).(data.MasterQ).New()
}

func CtxVerifierConfig(entry *config.VerifierConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, verifierConfigKey, entry)
	}
}

func VerifierConfig(r *http.Request) *config.VerifierConfig {
	return r.Context().Value(verifierConfigKey).(*config.VerifierConfig)
}

func CtxStateContract(entry *stateabi.State) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, stateContractKey, entry)
	}
}

func StateContract(r *http.Request) *stateabi.State {
	return r.Context().Value(stateContractKey).(*stateabi.State)
}

func CtxIssuer(iss *issuer.Issuer) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, issuerCtxKey, iss)
	}
}

func Issuer(r *http.Request) *issuer.Issuer {
	return r.Context().Value(issuerCtxKey).(*issuer.Issuer)
}

func CtxVaultClient(vaultClient *vault.VaultClient) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, vaultClientCtxKey, vaultClient)
	}
}

func VaultClient(r *http.Request) *vault.VaultClient {
	return r.Context().Value(vaultClientCtxKey).(*vault.VaultClient)
}

func CtxEthClient(client *ethclient.Client) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ethClientCtxKey, client)
	}
}

func EthClient(r *http.Request) *ethclient.Client {
	return r.Context().Value(ethClientCtxKey).(*ethclient.Client)
}

func CtxPoints(pointsCon *points.Client) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, pointsCtxKey, pointsCon)
	}
}

func Points(r *http.Request) *points.Client {
	return r.Context().Value(pointsCtxKey).(*points.Client)
}
