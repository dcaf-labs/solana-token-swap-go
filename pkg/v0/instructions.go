// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package spl_token_swap

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "SplTokenSwap"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_Swap = ag_binary.TypeID([8]byte{248, 198, 158, 145, 225, 117, 135, 200})

	Instruction_DepositAllTokenTypes = ag_binary.TypeID([8]byte{32, 95, 69, 60, 75, 79, 205, 238})

	Instruction_WithdrawAllTokenTypes = ag_binary.TypeID([8]byte{189, 254, 156, 174, 210, 9, 164, 216})

	Instruction_DepositSingleTokenTypeExactAmountIn = ag_binary.TypeID([8]byte{99, 207, 4, 42, 88, 157, 45, 55})

	Instruction_WithdrawSingleTokenTypeExactAmountOut = ag_binary.TypeID([8]byte{243, 192, 70, 139, 209, 156, 8, 139})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_Swap:
		return "Swap"
	case Instruction_DepositAllTokenTypes:
		return "DepositAllTokenTypes"
	case Instruction_WithdrawAllTokenTypes:
		return "WithdrawAllTokenTypes"
	case Instruction_DepositSingleTokenTypeExactAmountIn:
		return "DepositSingleTokenTypeExactAmountIn"
	case Instruction_WithdrawSingleTokenTypeExactAmountOut:
		return "WithdrawSingleTokenTypeExactAmountOut"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"initialize", (*Initialize)(nil),
		},
		{
			"swap", (*Swap)(nil),
		},
		{
			"deposit_all_token_types", (*DepositAllTokenTypes)(nil),
		},
		{
			"withdraw_all_token_types", (*WithdrawAllTokenTypes)(nil),
		},
		{
			"deposit_single_token_type_exact_amount_in", (*DepositSingleTokenTypeExactAmountIn)(nil),
		},
		{
			"withdraw_single_token_type_exact_amount_out", (*WithdrawSingleTokenTypeExactAmountOut)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBinEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBinDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
