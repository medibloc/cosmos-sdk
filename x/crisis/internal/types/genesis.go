package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - crisis genesis state
type GenesisState struct {
	ConstantFee sdk.Coin `json:"constant_fee" yaml:"constant_fee"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(constantFee sdk.Coin) GenesisState {
	return GenesisState{
		ConstantFee: constantFee,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{
		ConstantFee: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000000000000)), // Spend 1,000,000 MED for invariants check
	}
}

// ValidateGenesis - validate crisis genesis data
func ValidateGenesis(data GenesisState) error {
	if !data.ConstantFee.IsPositive() {
		return fmt.Errorf("constant fee must be positive: %s", data.ConstantFee)
	}
	return nil
}
