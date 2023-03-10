// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package spl_token_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DepositAllTokenTypes is the `depositAllTokenTypes` instruction.
type DepositAllTokenTypes struct {
	PoolTokenAmount     *uint64
	MaximumTokenAAmount *uint64
	MaximumTokenBAmount *uint64

	// [0] = [] swap
	//
	// [1] = [] authority
	//
	// [2] = [SIGNER] userTransferAuthority
	//
	// [3] = [WRITE] depositTokenA
	//
	// [4] = [WRITE] depositTokenB
	//
	// [5] = [WRITE] swapTokenA
	//
	// [6] = [WRITE] swapTokenB
	//
	// [7] = [WRITE] poolMint
	//
	// [8] = [WRITE] destination
	//
	// [9] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

type DepositAllTokenTypesAccounts struct {
	Swap                  ag_solanago.PublicKey
	Authority             ag_solanago.PublicKey
	UserTransferAuthority ag_solanago.PublicKey
	DepositTokenA         ag_solanago.PublicKey
	DepositTokenB         ag_solanago.PublicKey
	SwapTokenA            ag_solanago.PublicKey
	SwapTokenB            ag_solanago.PublicKey
	PoolMint              ag_solanago.PublicKey
	Destination           ag_solanago.PublicKey
	TokenProgram          ag_solanago.PublicKey
}

// NewDepositAllTokenTypesInstructionBuilder creates a new `DepositAllTokenTypes` instruction builder.
func NewDepositAllTokenTypesInstructionBuilder() *DepositAllTokenTypes {
	nd := &DepositAllTokenTypes{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

func (inst *DepositAllTokenTypes) GetDepositAllTokenTypesAccounts() *DepositAllTokenTypesAccounts {
	res := &DepositAllTokenTypesAccounts{}
	if inst.AccountMetaSlice[0] != nil {
		res.Swap = inst.AccountMetaSlice[0].PublicKey
	}
	if inst.AccountMetaSlice[1] != nil {
		res.Authority = inst.AccountMetaSlice[1].PublicKey
	}
	if inst.AccountMetaSlice[2] != nil {
		res.UserTransferAuthority = inst.AccountMetaSlice[2].PublicKey
	}
	if inst.AccountMetaSlice[3] != nil {
		res.DepositTokenA = inst.AccountMetaSlice[3].PublicKey
	}
	if inst.AccountMetaSlice[4] != nil {
		res.DepositTokenB = inst.AccountMetaSlice[4].PublicKey
	}
	if inst.AccountMetaSlice[5] != nil {
		res.SwapTokenA = inst.AccountMetaSlice[5].PublicKey
	}
	if inst.AccountMetaSlice[6] != nil {
		res.SwapTokenB = inst.AccountMetaSlice[6].PublicKey
	}
	if inst.AccountMetaSlice[7] != nil {
		res.PoolMint = inst.AccountMetaSlice[7].PublicKey
	}
	if inst.AccountMetaSlice[8] != nil {
		res.Destination = inst.AccountMetaSlice[8].PublicKey
	}
	if inst.AccountMetaSlice[9] != nil {
		res.TokenProgram = inst.AccountMetaSlice[9].PublicKey
	}
	return res
}

// SetPoolTokenAmount sets the "poolTokenAmount" parameter.
func (inst *DepositAllTokenTypes) SetPoolTokenAmount(poolTokenAmount uint64) *DepositAllTokenTypes {
	inst.PoolTokenAmount = &poolTokenAmount
	return inst
}

// SetMaximumTokenAAmount sets the "maximumTokenAAmount" parameter.
func (inst *DepositAllTokenTypes) SetMaximumTokenAAmount(maximumTokenAAmount uint64) *DepositAllTokenTypes {
	inst.MaximumTokenAAmount = &maximumTokenAAmount
	return inst
}

// SetMaximumTokenBAmount sets the "maximumTokenBAmount" parameter.
func (inst *DepositAllTokenTypes) SetMaximumTokenBAmount(maximumTokenBAmount uint64) *DepositAllTokenTypes {
	inst.MaximumTokenBAmount = &maximumTokenBAmount
	return inst
}

// SetSwapAccount sets the "swap" account.
func (inst *DepositAllTokenTypes) SetSwapAccount(swap ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(swap)
	return inst
}

// GetSwapAccount gets the "swap" account.
func (inst *DepositAllTokenTypes) GetSwapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *DepositAllTokenTypes) SetAuthorityAccount(authority ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *DepositAllTokenTypes) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserTransferAuthorityAccount sets the "userTransferAuthority" account.
func (inst *DepositAllTokenTypes) SetUserTransferAuthorityAccount(userTransferAuthority ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userTransferAuthority).SIGNER()
	return inst
}

// GetUserTransferAuthorityAccount gets the "userTransferAuthority" account.
func (inst *DepositAllTokenTypes) GetUserTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetDepositTokenAAccount sets the "depositTokenA" account.
func (inst *DepositAllTokenTypes) SetDepositTokenAAccount(depositTokenA ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(depositTokenA).WRITE()
	return inst
}

// GetDepositTokenAAccount gets the "depositTokenA" account.
func (inst *DepositAllTokenTypes) GetDepositTokenAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetDepositTokenBAccount sets the "depositTokenB" account.
func (inst *DepositAllTokenTypes) SetDepositTokenBAccount(depositTokenB ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(depositTokenB).WRITE()
	return inst
}

// GetDepositTokenBAccount gets the "depositTokenB" account.
func (inst *DepositAllTokenTypes) GetDepositTokenBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSwapTokenAAccount sets the "swapTokenA" account.
func (inst *DepositAllTokenTypes) SetSwapTokenAAccount(swapTokenA ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(swapTokenA).WRITE()
	return inst
}

// GetSwapTokenAAccount gets the "swapTokenA" account.
func (inst *DepositAllTokenTypes) GetSwapTokenAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSwapTokenBAccount sets the "swapTokenB" account.
func (inst *DepositAllTokenTypes) SetSwapTokenBAccount(swapTokenB ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(swapTokenB).WRITE()
	return inst
}

// GetSwapTokenBAccount gets the "swapTokenB" account.
func (inst *DepositAllTokenTypes) GetSwapTokenBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPoolMintAccount sets the "poolMint" account.
func (inst *DepositAllTokenTypes) SetPoolMintAccount(poolMint ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(poolMint).WRITE()
	return inst
}

// GetPoolMintAccount gets the "poolMint" account.
func (inst *DepositAllTokenTypes) GetPoolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetDestinationAccount sets the "destination" account.
func (inst *DepositAllTokenTypes) SetDestinationAccount(destination ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationAccount gets the "destination" account.
func (inst *DepositAllTokenTypes) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *DepositAllTokenTypes) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DepositAllTokenTypes {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *DepositAllTokenTypes) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst DepositAllTokenTypes) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DepositAllTokenTypes,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DepositAllTokenTypes) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DepositAllTokenTypes) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PoolTokenAmount == nil {
			return errors.New("PoolTokenAmount parameter is not set")
		}
		if inst.MaximumTokenAAmount == nil {
			return errors.New("MaximumTokenAAmount parameter is not set")
		}
		if inst.MaximumTokenBAmount == nil {
			return errors.New("MaximumTokenBAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Swap is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UserTransferAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.DepositTokenA is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.DepositTokenB is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SwapTokenA is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SwapTokenB is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.PoolMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Destination is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *DepositAllTokenTypes) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DepositAllTokenTypes")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    PoolTokenAmount", *inst.PoolTokenAmount))
						paramsBranch.Child(ag_format.Param("MaximumTokenAAmount", *inst.MaximumTokenAAmount))
						paramsBranch.Child(ag_format.Param("MaximumTokenBAmount", *inst.MaximumTokenBAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 swap", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("userTransferAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        depositTokenA", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        depositTokenB", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           swapTokenA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           swapTokenB", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("             poolMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          destination", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj DepositAllTokenTypes) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PoolTokenAmount` param:
	err = encoder.Encode(obj.PoolTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `MaximumTokenAAmount` param:
	err = encoder.Encode(obj.MaximumTokenAAmount)
	if err != nil {
		return err
	}
	// Serialize `MaximumTokenBAmount` param:
	err = encoder.Encode(obj.MaximumTokenBAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DepositAllTokenTypes) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PoolTokenAmount`:
	err = decoder.Decode(&obj.PoolTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `MaximumTokenAAmount`:
	err = decoder.Decode(&obj.MaximumTokenAAmount)
	if err != nil {
		return err
	}
	// Deserialize `MaximumTokenBAmount`:
	err = decoder.Decode(&obj.MaximumTokenBAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewDepositAllTokenTypesInstruction declares a new DepositAllTokenTypes instruction with the provided parameters and accounts.
func NewDepositAllTokenTypesInstruction(
	// Parameters:
	poolTokenAmount uint64,
	maximumTokenAAmount uint64,
	maximumTokenBAmount uint64,
	// Accounts:
	swap ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	userTransferAuthority ag_solanago.PublicKey,
	depositTokenA ag_solanago.PublicKey,
	depositTokenB ag_solanago.PublicKey,
	swapTokenA ag_solanago.PublicKey,
	swapTokenB ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	destination ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *DepositAllTokenTypes {
	return NewDepositAllTokenTypesInstructionBuilder().
		SetPoolTokenAmount(poolTokenAmount).
		SetMaximumTokenAAmount(maximumTokenAAmount).
		SetMaximumTokenBAmount(maximumTokenBAmount).
		SetSwapAccount(swap).
		SetAuthorityAccount(authority).
		SetUserTransferAuthorityAccount(userTransferAuthority).
		SetDepositTokenAAccount(depositTokenA).
		SetDepositTokenBAccount(depositTokenB).
		SetSwapTokenAAccount(swapTokenA).
		SetSwapTokenBAccount(swapTokenB).
		SetPoolMintAccount(poolMint).
		SetDestinationAccount(destination).
		SetTokenProgramAccount(tokenProgram)
}
