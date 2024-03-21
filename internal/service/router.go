package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	stateabi "github.com/iden3/contracts-abi/state/go/abi"
	"github.com/rarimo/rarime-passport-verifier/internal/data/pg"
	"github.com/rarimo/rarime-passport-verifier/internal/service/api/handlers"
	"github.com/rarimo/rarime-passport-verifier/internal/service/issuer"
	"github.com/rarimo/rarime-passport-verifier/internal/service/vault"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	ethCli, err := ethclient.Dial(s.cfg.NetworkConfig().EthRPC)
	if err != nil {
		s.log.WithError(err).Fatal("failed to dial connect via Ethereum RPC")
	}

	stateContract, err := stateabi.NewState(common.HexToAddress(s.cfg.NetworkConfig().StateContract), ethCli)
	if err != nil {
		s.log.WithError(err).Fatal("failed to init state contract")
	}

	vaultClient, err := vault.NewVaultClient(s.cfg.VaultConfig())
	if err != nil {
		s.log.WithError(err).Fatal("failed to init new vault client")
	}

	issuerLogin, issuerPassword, err := vaultClient.IssuerAuthData()
	if err != nil {
		s.log.WithError(err).Fatal("failed to get issuer auth data from the vault")
	}

	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxMasterQ(pg.NewMasterQ(s.cfg.DB())),
			handlers.CtxVerifierConfig(s.cfg.VerifierConfig()),
			handlers.CtxStateContract(stateContract),
			handlers.CtxIssuer(issuer.New(
				s.cfg.Log().WithField("service", "issuer"),
				s.cfg.IssuerConfig(),
				issuerLogin, issuerPassword,
			)),
			handlers.CtxVaultClient(vaultClient),
			handlers.CtxEthClient(ethCli),
			handlers.CtxPoints(s.cfg.Points()),
		),
	)
	r.Route("/integrations/identity-provider-service", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/create-identity", handlers.CreateIdentity)
			r.Get("/gist-data", handlers.GetGistData)
		})
	})

	return r
}
