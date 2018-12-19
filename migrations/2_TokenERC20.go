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

type TokenERC20Deployer struct{}

func (d *TokenERC20Deployer) Deploy(ctx context.Context, network *network.Network) (common.Address, *types.Transaction, interface{}, error) {
	account := network.Accounts()[0]
	network.UnlockWithPrompt(account)

	auth := network.NewTransactor(account)
	address, transaction, contract, err := bindings.DeployTokenERC20(auth, network.Client())
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	session := &bindings.TokenERC20Session{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *TokenERC20Deployer) Bind(ctx context.Context, network *network.Network, address common.Address) (interface{}, error) {
	account := network.Accounts()[0]
	network.UnlockWithPrompt(account)

	auth := network.NewTransactor(account)
	contract, err := bindings.NewTokenERC20(address, network.Client())
	if err != nil {
		return nil, err
	}

	session := &bindings.TokenERC20Session{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("TokenERC20", &TokenERC20Deployer{})

	migration.AddMigration(&migration.Migration{
		Number: 2,
		F: func(ctx context.Context, network *network.Network) error {
			if err := contract.Deploy(ctx, "TokenERC20", network); err != nil {
				return err
			}

			return nil
		},
	})
}
