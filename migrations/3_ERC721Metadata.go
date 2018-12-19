package migrations

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"

	"perigord/bindings"
)

type ERC721MetadataDeployer struct{}

func (d *ERC721MetadataDeployer) Deploy(ctx context.Context, network *network.Network) (common.Address, *types.Transaction, interface{}, error) {
	account := network.Accounts()[0]
	network.UnlockWithPrompt(account)

	auth := network.NewTransactor(account)
	address, transaction, contract, err := bindings.DeployERC721Metadata(auth, network.Client())
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	session := &bindings.ERC721MetadataSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *ERC721MetadataDeployer) Bind(ctx context.Context, network *network.Network, address common.Address) (interface{}, error) {
	account := network.Accounts()[0]
	network.UnlockWithPrompt(account)

	auth := network.NewTransactor(account)
	contract, err := bindings.NewERC721Metadata(address, network.Client())
	if err != nil {
		return nil, err
	}

	session := &bindings.ERC721MetadataSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("ERC721Metadata", &ERC721MetadataDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 3,
		F: func(ctx context.Context, network *network.Network) error {
			if err := contract.Deploy(ctx, "ERC721Metadata", network); err != nil {
				return err
			}

			return nil
		},
	})
}
