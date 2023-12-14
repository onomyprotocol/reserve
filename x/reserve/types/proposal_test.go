package types_test

import (
	"reserve/testutil/sample"
	"reserve/x/reserve/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func TestCreateDenomProposal_ValidateBasic(t *testing.T) { //nolint:dupl // test template

	type fields struct {
		Sender      string
		Title       string
		Description string
		Metadata    banktypes.Metadata
		Rate        []sdk.Uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "positive",
			fields: fields{
				Sender:      sample.AccAddress(),
				Title:       "title",
				Description: "desc",
				Metadata:    banktypes.Metadata{},
				Rate:        []sdk.Uint{sdk.NewUintFromString("1"), sdk.NewUintFromString("1")},
			},
		},
		{
			name: "negative_invalid_sender",
			fields: fields{
				Sender:      "invalid-sender",
				Title:       "title",
				Description: "desc",
				Metadata:    banktypes.Metadata{},
				Rate:        []sdk.Uint{sdk.NewUintFromString("1"), sdk.NewUintFromString("1")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &types.CreateDenomProposal{
				Sender:      tt.fields.Sender,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Metadata:    &tt.fields.Metadata,
				Rate:        tt.fields.Rate,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
